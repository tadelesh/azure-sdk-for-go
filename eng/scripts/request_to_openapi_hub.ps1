param([string]$specRoot)

function executeSingleRequest($configFile)
{
    $content = Get-Content $configFile -Raw
    if ($content -notmatch "``````\s*yaml\s+\`$\(go\).+\`$\(track2\)") {
        Write-Host "No go track2 config found, skip $configFile"
        return
    }

    if ($content -match "module-name:\s+sdk/resourcemanager/(?<rpName>.*)/(?<packageName>.*)\r\n") {
        $rpName = $matches["rpName"]
        $packageName = $matches["packageName"]

        $url = "https://codegenappdev.azurewebsites.net/codegenerations/go-$rpName"
        
        $body = ConvertTo-Json @{
            resourceProvider = "$rpName"
            sdk              = "go"
            type             = "CI"
            codegenRepo      = @{
                type   = "github"
                path   = "https://github.com/Azure/depth-coverage-pipeline"
                branch = "go_pipeline"
            }
            sdkRepo          = @{
                type   = "github"
                path   = "https://github.com/tadelesh/azure-sdk-for-go"
                branch = "main"
            }
            swaggerRepo      = @{
                type   = "github"
                path   = "https://github.com/tadelesh/azure-rest-api-specs"
                branch = "go_track2_test"
            }
        }

        $headers = @{
            "Content-Type"  = "application/json"
        }

        Invoke-RestMethod -Uri $url -Body $body -Headers $headers -Method "PUT"
        
        # $url = "https://codegenappdev.azurewebsites.net/codegenerations/go-$rp"
        # Invoke-RestMethod -Uri $url -Body $body -Headers $headers -Method "DELETE"

        # $url = "https://codegenappdev.azurewebsites.net/codegenerations/go-$rp/detail"
        # Invoke-RestMethod -Uri $url -Body $body -Headers $headers -Method "GET"
    }
}

Get-ChildItem -recurse -path $specRoot -filter readme.go.md | ForEach-Object {
    if ($_ -match "specification.*resource-manager.*readme.go.md") {
        Write-Host "Processing $_ ..."
        executeSingleRequest $_
    }
}

