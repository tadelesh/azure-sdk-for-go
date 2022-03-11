Param(
    [string] $sdkPath,
    [string] $specPath,
    [bool] $generateSDK = $true,
    [bool] $generateTest = $false,
    [bool] $executeMockTest = $false
)

$AUTOREST_TEST_PACKAGE_URL = "D:\Workspace\Azure\azure-sdk-tools\tools\sdk-testgen\packages\autorest.gotest"
$AUTOREST_GO_VERSION = "@autorest/go@4.0.0-preview.37"
$AUTOREST_CONFIG_FILE = "autorest.md"
$AUTOREST_CORE_VERSION = "3.6.2"

function executeSingleGenerate($readmePath, $sepcRPName)
{
    Write-Host "Processing Spec RP $sepcRPName ..."
    $content = Get-Content $readmePath -Raw
    if ($content -notmatch "\`$\(track2\)")
    {
        Write-Host "No GO track2 config, skip ..."
        return
        $rpName = $sepcRPName -replace "-", ""
        $content = $content.Replace("`````` yaml `$(go)", "`````` yaml `$(go) && !`$(track2)")
        $content += "
``````yaml `$(go) && `$(track2)
license-header: MICROSOFT_MIT_NO_VERSION
module-name: sdk/resourcemanager/$rpName/arm$rpName
module: github.com/Azure/azure-sdk-for-go/`$(module-name)
output-folder: `$(go-sdk-folder)/`$(module-name)
azure-arm: true
modelerfour:
  lenient-model-deduplication: true
``````
"
        Set-Content (Join-Path $specPath "specification" $sepcRPName "resource-manager" "readme.go.md") $content
        $packageName = "arm$rpName"
    }
    else
    {
        $content -match "module-name: sdk\/resourcemanager\/(?<rpName>.*)\/(?<packageName>.*)"
        $rpName = $matches["rpName"]
        $packageName = $matches["packageName"] -replace "`n", "" -replace "`r", ""
    }
    
    if ($generateSDK)
    {
        Write-Host "Generate sdk code for RP $rpName with Package $packageName ..."
        Write-Host "generator release-v2 $sdkPath $specPath $rpName $packageName --spec-rp-name=$sepcRPName --skip-create-branch=true"
        generator release-v2 $sdkPath $specPath $rpName $packageName --spec-rp-name=$sepcRPName --skip-create-branch=true --skip-generate-example=true
        if ($LASTEXITCODE)
        {
            Write-Host "##[error] generate sdk code error for RP $rpName with Package $packageName"
            return
        }
        Set-Location (Join-Path $sdkPath "sdk" "resourcemanager" $rpName $packageName)
        if (-not (checkResult "sdk code" $rpName $packageName))
        {
            return
        }
    }

    if ($generateTest)
    {
        Write-Host "Generate test code for RP $rpName with Package $packageName ..."
        Set-Location (Join-Path $sdkPath "sdk" "resourcemanager" $rpName $packageName)
        Write-Host "autorest --version=$AUTOREST_CORE_VERSION --use=$AUTOREST_GO_VERSION --use=$AUTOREST_TEST_PACKAGE_URL --file-prefix=zz_generated_ --track2 --go --output-folder=. --clear-output-folder=false --testmodeler.generate-mock-test --testmodeler.generate-sdk-example --testmodeler.generate-scenario-test --generate-sdk=false .\$AUTOREST_CONFIG_FILE"
        autorest --version=$AUTOREST_CORE_VERSION --use=$AUTOREST_GO_VERSION --use=D:\Workspace\Azure\azure-sdk-tools\tools\sdk-testgen\packages\autorest.testmodeler --use=$AUTOREST_TEST_PACKAGE_URL --file-prefix=zz_generated_ --track2 --go --output-folder=. --clear-output-folder=false --testmodeler.generate-mock-test --testmodeler.generate-sdk-example --testmodeler.generate-scenario-test --generate-sdk=true .\$AUTOREST_CONFIG_FILE
        if ($LASTEXITCODE)
        {
            Write-Host "##[error] generation test code error for RP $rpName with Package $packageName"
            return
        }
        if (-not (checkResult "test code" $rpName $packageName))
        {
            return
        }
    }

    if($executeMockTest)
    {
        Write-Host "Execute mock test for RP $rpName with Package $packageName ..."
        Set-Location (Join-Path $sdkPath "sdk" "resourcemanager" $rpName $packageName)
        $Env:AZURE_TENANT_ID = "mock-test"
        $Env:AZURE_CLIENT_ID = "mock-test"
        $Env:AZURE_USERNAME = "mock-test"
        $Env:AZURE_PASSWORD = "mock-test"
        go test -v -coverprofile coverage.txt
        if ($LASTEXITCODE)
        {
            Write-Host "##[error] execute mock test error for RP $rpName with Package $packageName"
            return
        }
    }
}

function checkResult($checkName, $rpName, $packageName)
{
    gofmt -s -w .
    go mod tidy
    if ($LASTEXITCODE)
    {
        Write-Host "##[error] mod tidy error for $checkName of RP $rpName with Package $packageName"
        return $false
    }
    go build .
    if ($LASTEXITCODE)
    {
        Write-Host "##[error] build error for $checkName of RP $rpName with Package $packageName"
        return $false
    }
    go vet .
    if ($LASTEXITCODE)
    {
        Write-Host "##[error] vet error for $checkName of RP $rpName with Package $packageName"
        return $false
    }
    return $true
}

Get-ChildItem -recurse -path (Join-Path $specPath "specification") -filter readme.go.md | ForEach-Object {
    if ($_.FullName -match "[\/|\\]specification[\/|\\](?<specRPName>.*)[\/|\\]resource-manager[\/|\\]readme.go.md")
    {
        executeSingleGenerate $_ $matches["specRPName"]
    }
}
