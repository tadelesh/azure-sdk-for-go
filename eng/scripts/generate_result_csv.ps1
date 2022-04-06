Param(
    [string] $sdkPath,
    [string] $specPath,
    [string] $logPath
)

function getLogResults($sdkPath, $specPath, $logPath)
{
    $logContent = Get-Content $logPath
    $lineNo = 0
    $resultArray = @()
    $resultObject = $null
    while ($lineNo -lt $logContent.Length)
    {
        $logLine = $logContent[$lineNo]
        if ($logLine -match "Processing Spec RP (?<specRPName>.+) ...")
        {
            if ($null -ne $resultObject)
            {
                $resultArray += $resultObject
            }
            $specRPName = $matches["specRPName"]
            Write-Host "Start to process log for RP: $specRPName"
            $resultObject = [PSCustomObject]@{
                specRPName          = $specRPName
                rpName              = ""
                packageName         = ""
                defaultTag          = ""
                previewTags         = ""
                stableTags          = ""
                latestPreviewTag    = ""
                latestStableTag     = ""
                directives          = ""
                sdkAutorestFailure  = ""
                generateToolFailure = ""
                sdkCodeBuild        = ""
                sdkGenerationError  = ""
                testAutorestFailure = ""
                testCodeBuild       = ""
                testGenerationError = ""
                listOfTestCase      = ""
                listOfPassCase      = ""
                listOfSkipCase      = ""
                listOfFailCase      = ""
                listOfFailExample   = ""
                listOfFailCaseType  = ""
                numOfTestCase       = 0
                numOfPassCase       = 0
                numOfFailCase       = 0
                numOfSkipCase       = 0
                passRate            = 0.0
                failRate            = 0.0
                skipRate            = 0.0
                codeCoverage        = 0.0
                testExecutionError  = ""
            }
            $shouldGetBasicInfo = $true
            $lineNo = $lineNo + 1
        }
        elseif ($logLine -match "Generate sdk code for RP (?<rpName>.+) with Package (?<packageName>.+) ...")
        {
            $rpName = $matches["rpName"]
            $packageName = $matches["packageName"]
            Write-Host "Processing SDK code generation result log for RP: $rpName Package: $packageName"
            if ($shouldGetBasicInfo)
            {
                getBasicInfo $resultObject $sdkPath $specPath $rpName $packageName
                $shouldGetBasicInfo = $false
            }

            $lineNo = $lineNo + 1
            $errorDetail = ""
            while ($lineNo -lt $logContent.Length)
            {
                $logLine = $logContent[$lineNo]
                if ($logLine -match "^fatal\s+\|")
                {
                    if ($errorDetail -ne "")
                    {
                        $errorDetail = "$errorDetail`n$logLine"
                    }
                    else
                    {
                        $errorDetail = $logLine
                    }
                }
                elseif ($logLine -match "^\s+Error:")
                {
                    $resultObject.sdkAutorestFailure = $errorDetail
                    $errorDetail = ""
                }
                elseif ($logLine -match "^failed to finish release generation process")
                {
                    $resultObject.generateToolFailure = "detail to be added"
                }
                elseif ($logLine -match "^##\[error\] mod tidy error for")
                {
                    $resultObject.sdkCodeBuild = "fmt/modtidy"
                }
                elseif ($logLine -match "^##\[error\] build error for")
                {
                    $resultObject.sdkCodeBuild = "build"
                }
                elseif ($logLine -match "^##\[error\] vet error for")
                {
                    $resultObject.sdkCodeBuild = "vet"
                }
                elseif ($logLine -match "^##\[error\] generate sdk code error for")
                {
                    $resultObject.sdkGenerationError = "×"
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
            $rpName = $matches["rpName"]
            $packageName = $matches["packageName"]
            Write-Host "Processing test code generation result log for RP: $rpName Package: $packageName"
            if ($shouldGetBasicInfo)
            {
                getBasicInfo $resultObject $sdkPath $specPath $rpName $packageName
                $shouldGetBasicInfo = $false
            }
            
            $lineNo = $lineNo + 1
            $errorDetail = ""
            while ($lineNo -lt $logContent.Length)
            {
                $logLine = $logContent[$lineNo]
                if ($logLine -match "^fatal\s+\|")
                {
                    if ($errorDetail -ne "")
                    {
                        $errorDetail = "$errorDetail`n$logLine"
                    }
                    else
                    {
                        $errorDetail = $logLine
                    }
                }
                elseif ($logLine -match "^\s+Error:")
                {
                    $resultObject.testAutorestFailure = $errorDetail
                    $errorDetail = ""
                }
                elseif ($logLine -match "^##\[error\] mod tidy error for")
                {
                    $resultObject.testCodeBuild = "fmt/modtidy"
                }
                elseif ($logLine -match "^##\[error\] build error for")
                {
                    $resultObject.testCodeBuild = "build"
                }
                elseif ($logLine -match "^##\[error\] vet error for")
                {
                    $resultObject.testCodeBuild = "vet"
                }
                elseif ($logLine -match "^##\[error\] generation test code error")
                {
                    $resultObject.testGenerationError = "×"
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
            $rpName = $matches["rpName"]
            $packageName = $matches["packageName"]
            Write-Host "Processing mock test result log for RP: $rpName Package: $packageName"
            if ($shouldGetBasicInfo)
            {
                getBasicInfo $resultObject $sdkPath $specPath $rpName $packageName
                $shouldGetBasicInfo = $false
            }
            
            $lineNo = $lineNo + 1
            $errorDetail = ""
            $listOfTestCase = @()
            $listOfPassCase = @()
            $listOfSkipCase = @()
            $listOfFailCase = @()
            $listOfFailExample = @()
            $listOfFailCaseType = @()
            while ($lineNo -lt $logContent.Length)
            {
                $logLine = $logContent[$lineNo]
                if ($logLine -match "^=== RUN\s+(?<testcaseName>.+)")
                {
                    $listOfTestCase += $matches["testcaseName"]
                    $errorDetail = ""
                }
                elseif ($logLine -match "^--- FAIL:\s+(?<testcaseName>.+)\s+")
                {
                    $listOfFailCase += $matches["testcaseName"]
                    if ($errorDetail -match "`"INVALID_FORMAT\\`"")
                    {
                        $listOfFailCaseType += "Invalid fomat"
                    }
                    elseif ($errorDetail -match "`"OBJECT_MISSING_REQUIRED_PROPERTY\\`"")
                    {
                        $listOfFailCaseType += "Missing required property"
                    }
                    elseif ($errorDetail -match "`"INVALID_TYPE\\`"")
                    {
                        $listOfFailCaseType += "Invalid type"
                    }
                    elseif ($errorDetail -match "`"OPERATION_NOT_FOUND_IN_CACHE_WITH_API\\`"")
                    {
                        $listOfFailCaseType += "Operation not found with api"
                    }
                    elseif ($errorDetail -match "`"OPERATION_NOT_FOUND_IN_CACHE_WITH_PROVIDER\\`"")
                    {
                        $listOfFailCaseType += "Operation not found with provider"
                    }
                    elseif ($errorDetail -match "`"OPERATION_NOT_FOUND_IN_CACHE\\`"")
                    {
                        $listOfFailCaseType += "Operation not found"
                    }
                    elseif ($errorDetail -match "`"DISCRIMINATOR_VALUE_NOT_FOUND\\`"")
                    {
                        $listOfFailCaseType += "Discriminator value not found"
                    }
                    elseif ($errorDetail -match "`"INVALID_CONTENT_TYPE\\`"")
                    {
                        $listOfFailCaseType += "Invalid content type"
                    }
                    elseif ($errorDetail -match "Mock response is not equal to example response for example")
                    {
                        $listOfFailCaseType += "Example response not match with schema"
                    }
                    elseif ($errorDetail -match "Example Not Match")
                    {
                        $listOfFailCaseType += "Example request not match with schema"
                    }
                    elseif ($errorDetail -match "Example Not Found")
                    {
                        $listOfFailCaseType += "Example Not Found (mainly because the naming duplication)"
                    }
                    elseif ($errorDetail -match "Can't find a GET operation nearby")
                    {
                        $listOfFailCaseType += "Can not get cooresponding get operation for LRO"
                    }
                    elseif ($errorDetail -match "the response did not contain a body")
                    {
                        $listOfFailCaseType += "LRO response not contains 200 response"
                    }
                    elseif ($errorDetail -match "Wrong response example for operation")
                    {
                        $listOfFailCaseType += "Wrong response example for operation"
                    }
                    elseif ($errorDetail -match "parsing time")
                    {
                        $listOfFailCaseType += "Response has wrong datetime value"
                    }
                    elseif ($errorDetail -match "json: cannot unmarshal")
                    {
                        if ($errorDetail -match "type map\[string\]interface \{\}")
                        {
                            $listOfFailCaseType += "Object additional property with wrong example value"
                        }
                        else
                        {
                            $listOfFailCaseType += "Example response value not match with schema"
                        }
                    }
                    elseif ($errorDetail -match "unmarshalling type")
                    {
                        $listOfFailCaseType += "Example response value not match with schema"
                    }
                    else
                    {
                        $listOfFailCaseType += "N/A"
                    }
                    if ($errorDetail -match "for example (?<exampleFile>specification/[^:]+):")
                    {
                        $listOfFailExample += $matches["exampleFile"]
                    }
                    else
                    {
                        $listOfFailExample += "N/A"
                    }
                }
                elseif ($logLine -match "^--- PASS:\s+(?<testcaseName>.+)\s+")
                {
                    $listOfPassCase += $matches["testcaseName"]
                }
                elseif ($logLine -match "^--- SKIP:\s+(?<testcaseName>.+)\s+")
                {
                    $listOfSkipCase += $matches["testcaseName"]
                }
                elseif ($logLine -match "^coverage: (?<coverageRate>.+)% of statements")
                {
                    $resultObject.codeCoverage = $matches["coverageRate"]
                }
                elseif ($logLine -match "^##\[error\] execute mock test error")
                {
                    $resultObject.testExecutionError = "×"
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

            $resultObject.listOfTestCase = $listOfTestCase -join "`n"
            $resultObject.listOfPassCase = $listOfPassCase -join "`n"
            $resultObject.listOfSkipCase = $listOfSkipCase -join "`n"
            $resultObject.listOfFailCase = $listOfFailCase -join "`n"
            $resultObject.listOfFailCaseType = $listOfFailCaseType -join "`n"
            $resultObject.listOfFailExample = $listOfFailExample -join "`n"
            $resultObject.numOfTestCase = $listOfTestCase.length
            $resultObject.numOfPassCase = $listOfPassCase.length
            $resultObject.numOfFailCase = $listOfFailCase.length
            $resultObject.numOfSkipCase = $listOfSkipCase.length
            if ($listOfTestCase.length -ne 0)
            {
                $resultObject.passRate = $resultObject.numOfPassCase / $resultObject.numOfTestCase
                $resultObject.failRate = $resultObject.numOfFailCase / $resultObject.numOfTestCase
                $resultObject.skipRate = $resultObject.numOfSkipCase / $resultObject.numOfTestCase
            }
        }
        else
        {
            $lineNo = $lineNo + 1
        } 
    }
    $resultArray += $resultObject

    if (Test-Path -Path .\result.csv)
    {
        Remove-Item -Path .\result.csv
    }
    $resultArray | Export-Csv -Path .\result.csv
}

function getBasicInfo($resultObject, $sdkPath, $specPath, $rpName, $packageName)
{
    $resultObject.rpName = $rpName
    $resultObject.packageName = $packageName

    $content = Get-Content (Join-Path $sdkPath "sdk" "resourcemanager" $rpName $packageName "autorest.md") -Raw
    if ($content -match "(?<directive>directive:(.|\s)+)``````")
    {
        $resultObject.directives = "autorest.md:`n$($matches["directive"])`n"
    }

    $content = Get-Content (Join-Path $specPath "specification" $resultObject.specRPName "resource-manager" "readme.go.md") -Raw
    if ($content -match "`````` yaml \$\(go\) && \$\(track2\)[^``]+(?<directive>directive:[^``]+)``````")
    {
        $resultObject.directives = "readme.go.md:`n$($matches["directive"])`n"
    }

    $content = Get-Content (Join-Path $specPath "specification" $resultObject.specRPName "resource-manager" "readme.md") -Raw

    if ($content -match "``````\s*yaml\s[^``]+tag:\s+(?<defaultTag>\S+)\s")
    {
        $resultObject.defaultTag = $matches["defaultTag"]
    }

    $results = ([regex]"### Tag: (?<tag>\S+)(\n|\r)+").Matches($content)

    $previewTags = @()
    $stableTags = @()
            
    foreach ($result in $results)
    {
        $tag = $result.Groups["tag"].Value
        if ($tag -match "preview")
        {
            $previewTags += $tag
        }
        else
        {
            $stableTags += $tag
        }
    }
    
    $previewTags = $previewTags | Sort-Object -Descending
    $stableTags = $stableTags | Sort-Object -Descending

    $resultObject.previewTags = $previewTags -join "`n"
    $resultObject.stableTags = $stableTags -join "`n"
    if ($previewTags.Count -eq 1)
    {
        $resultObject.latestPreviewTag = $previewTags
    }
    elseif ($previewTags.Count -gt 1)
    {
        $resultObject.latestPreviewTag = $previewTags[0]
    }
    if ($stableTags.Count -eq 1)
    {
        $resultObject.latestStableTag = $stableTags
    }
    elseif ($stableTags.Count -gt 1)
    {
        $resultObject.latestStableTag = $stableTags[0]
    }
}

getLogResults $sdkPath $specPath $logPath
