Param(
    [string] $logPath,
    [string] $validationPath
)

function getComparingResult($logPath, $validationPath)
{
    $excel = New-Object -Com Excel.Application
    $workbook = $excel.Workbooks.Open($validationPath)
    $sheet = $workbook.Sheets.Item("model-result-controlplane")
    $rowCount = $sheet.UsedRange.Rows.Count
    $rpName = ""
    $validationResult = @{}
    for ($i = 2; $i -le $rowcount; $i++)
    {
        if ($sheet.Rows.Item($i).Columns.Item(1).Text -ne "")
        {
            $rpName = $sheet.Rows.Item($i).Columns.Item(1).Text
            $validationResult[$rpName] = @{}
            $validationResult[$rpName]["exampleNotFound"] = $false
            $validationResult[$rpName]["problemList"] = @()
            $validationResult[$rpName]["problemReasonList"] = @()
        }
        if ($sheet.Rows.Item($i).Columns.Item(4).Text -eq "XMS_EXAMPLE_NOTFOUND_ERROR")
        {
            $validationResult[$rpName]["exampleNotFound"] = $true
        }
        else
        {
            $validationResult[$rpName]["problemList"] += $sheet.Rows.Item($i).Columns.Item(3).Text
            $validationResult[$rpName]["problemReasonList"] += $sheet.Rows.Item($i).Columns.Item(4).Text
        }
    }
    $workbook.close($true)

    $logContent = Get-Content $logPath
    $lineNo = 0
    $resultArray = @()
    while ($lineNo -lt $logContent.Length)
    {
        $logLine = $logContent[$lineNo]
        if ($logLine -match "Processing Spec RP (?<specRPName>.+) ...")
        {
            $specRPName = $matches["specRPName"]
            Write-Host "Start to process log for RP: $specRPName"
            
            $lineNo = $lineNo + 1
        }
        elseif ($logLine -match "Generate sdk code for RP (?<rpName>.+) with Package (?<packageName>.+) ...")
        {
            $lineNo = $lineNo + 1
            while ($lineNo -lt $logContent.Length)
            {
                $logLine = $logContent[$lineNo]
                if ($logLine -match "^##\[error\] generate sdk code error for")
                {
                    Write-Host "RP: $rpName failed to generate SDK"
                }
                elseif ($logLine -match "^Generate |^Processing Spec RP|^Execute mock ")
                {
                    break
                }
                $lineNo = $lineNo + 1
            }
        }
        elseif ($logLine -match "Generate test code for RP (?<rpName>.+) with Package (?<packageName>.+) ...")
        {
            $lineNo = $lineNo + 1
            while ($lineNo -lt $logContent.Length)
            {
                $logLine = $logContent[$lineNo]
                if ($logLine -match "^##\[error\] generation test code error")
                {
                    Write-Host "RP: $rpName failed to generate SDK test"
                }
                elseif ($logLine -match "^Generate |^Processing Spec RP|^Execute mock ")
                {
                    break
                }
                $lineNo = $lineNo + 1
            }
        }
        elseif ($logLine -match "Execute mock test for RP (?<rpName>.+) with Package (?<packageName>.+) ...")
        {
            $lineNo = $lineNo + 1
            $errorDetail = ""
            $skipCase = $false
            while ($lineNo -lt $logContent.Length)
            {
                $logLine = $logContent[$lineNo]
                if ($logLine -match "^=== RUN\s+(?<testcaseName>.+)")
                {
                    $errorDetail = ""
                }
                elseif ($logLine -match "^--- FAIL:\s+(?<testcaseName>.+)\s+")
                {
                    if ($errorDetail -match "`"INVALID_FORMAT\\`"")
                    {
                        $mock_reason = "Example request has invalid format value"
                    }
                    elseif ($errorDetail -match "`"OBJECT_ADDITIONAL_PROPERTIES\\`"")
                    {
                        $mock_reason = "Example request not match with schema"
                    }
                    elseif ($errorDetail -match "`"OBJECT_MISSING_REQUIRED_PROPERTY\\`"")
                    {
                        $mock_reason = "Example request miss required property"
                    }
                    elseif ($errorDetail -match "`"INVALID_TYPE\\`"")
                    {
                        $mock_reason = "Example request has invalid type value"
                    }
                    elseif ($errorDetail -match "`"OPERATION_NOT_FOUND_IN_CACHE_WITH_API\\`"")
                    {
                        $mock_reason = "Operation do not have api version param"
                    }
                    elseif ($errorDetail -match "`"OPERATION_NOT_FOUND_IN_CACHE_WITH_PROVIDER\\`"")
                    {
                        $mock_reason = "OPERATION_NOT_FOUND_IN_CACHE_WITH_PROVIDER"
                    }
                    elseif ($errorDetail -match "`"OPERATION_NOT_FOUND_IN_CACHE\\`"")
                    {
                        $mock_reason = "OPERATION_NOT_FOUND_IN_CACHE"
                    }
                    elseif ($errorDetail -match "`"DISCRIMINATOR_VALUE_NOT_FOUND\\`"")
                    {
                        $mock_reason = "Example request missing discriminator value"
                    }
                    elseif ($errorDetail -match "`"INVALID_CONTENT_TYPE\\`"")
                    {
                        $mock_reason = "Example request is in invalid content type"
                    }
                    elseif ($errorDetail -match "Mock response is not equal to example response for example")
                    {
                        $mock_reason = "Example response not match with schema"
                    }
                    elseif ($errorDetail -match "Example Not Match")
                    {
                        if ($errorDetail -match "query parameter api-version"){
                            $mock_reason = "Wrong operaion api version"
                        }else{
                            $mock_reason = "Example request not match with schema"
                        }
                    }
                    elseif ($errorDetail -match "Example Not Found")
                    {
                        $mock_reason = "Duplicated example id"
                    }
                    elseif ($errorDetail -match "Can't find a GET operation nearby")
                    {
                        $mock_reason = "LRO operation has no cooresponding get operation"
                    }
                    elseif ($errorDetail -match "the response did not contain a body")
                    {
                        $mock_reason = "LRO response not contains 200 response"
                    }
                    elseif ($errorDetail -match "parsing time")
                    {
                        $mock_reason = "Response has wrong datetime value"
                    }
                    elseif ($errorDetail -match "json: cannot unmarshal")
                    {
                        $mock_reason = "Example response not match with schema"
                    }
                    elseif ($errorDetail -match "unmarshalling type")
                    {
                        $mock_reason = "Example response not match with schema"
                    }
                    else
                    {
                        $mock_reason = "N/A"
                    }

                    if ($null -ne $validationResult[$specRPName])
                    {
                        $found = $false
                        for ($i = 0; $i -lt $validationResult[$specRPName]["problemList"].length; $i++)
                        {
                            if ($errorDetail -like "*$($validationResult[$specRPName]["problemList"][$i])*")
                            {
                                $found = $true
                                break
                            }
                        }
                        if ($found -eq $false)
                        {
                            if ($errorDetail -match "for example (?<exampleFile>specification/[^:]+):")
                            {
                                $errorExample = $matches["exampleFile"]
                            }
                            else
                            {
                                $errorExample += "N/A"
                            }
                    
                            $resultArray += [PSCustomObject]@{
                                specRPName                       = $specRPName
                                errorExample                     = $errorExample
                                errorExampleReasonFromValidation = ""
                                errorExampleReasonFromMock       = $mock_reason
                            }
                        }
                        else
                        {
                            $resultArray += [PSCustomObject]@{
                                specRPName                       = $specRPName
                                errorExample                     = $errorExample
                                errorExampleReasonFromValidation = $validationResult[$specRPName]["problemReasonList"][$i]
                                errorExampleReasonFromMock       = ""
                            }
                        }
                    }
                    else
                    {
                        if ($errorDetail -match "for example (?<exampleFile>specification/[^:]+):")
                        {
                            $errorExample = $matches["exampleFile"]
                        }
                        else
                        {
                            $errorExample += "N/A"
                        }

                        $resultArray += [PSCustomObject]@{
                            specRPName                       = $specRPName
                            errorExample                     = $errorExample
                            errorExampleReasonFromValidation = ""
                            errorExampleReasonFromMock       = $mock_reason
                        }
                    }
                }
                elseif ($logLine -match "^--- SKIP:\s+(?<testcaseName>.+)\s+")
                {
                    $skipCase = $true
                }
                elseif ($logLine -match "^Generate |^Processing Spec RP|^Execute mock ")
                {
                    break
                }
                else
                {
                    $errorDetail = "$errorDetail`n$logLine"
                }
                $lineNo = $lineNo + 1
            }

            if ($skipCase -eq $true)
            {
                if ($null -ne $validationResult[$specRPName])
                {
                    if ($validationResult[$specRPName]["exampleNotFound"] -eq $false)
                    {
                        $resultArray += [PSCustomObject]@{
                            specRPName                       = $specRPName
                            errorExample                     = ""
                            errorExampleReasonFromValidation = ""
                            errorExampleReasonFromMock       = "XMS_EXAMPLE_NOTFOUND_ERROR"
                        }
                    } else {
                        $resultArray += [PSCustomObject]@{
                            specRPName                       = $specRPName
                            errorExample                     = ""
                            errorExampleReasonFromValidation = "XMS_EXAMPLE_NOTFOUND_ERROR"
                            errorExampleReasonFromMock       = ""
                        }
                    }
                }
                else
                {
                    $resultArray += [PSCustomObject]@{
                        specRPName                       = $specRPName
                        errorExample                     = ""
                        errorExampleReasonFromValidation = ""
                        errorExampleReasonFromMock       = "XMS_EXAMPLE_NOTFOUND_ERROR"
                    }
                }
            }
        }
        else
        {
            $lineNo = $lineNo + 1
        } 
    }

    if (Test-Path -Path .\comparing_result.csv)
    {
        Remove-Item -Path .\comparing_result.csv
    }
    $resultArray | Export-Csv -Path .\comparing_result.csv
}

getComparingResult $logPath $validationPath