$IgnoreDirs = @(".git", "node_modules")
$IgnoreFiles = @("Thumbs.db", ".DS_Store")

function Show-Tree($path, $indent = "") {
    Get-ChildItem -LiteralPath $path -Force | ForEach-Object {
        if ($_.PSIsContainer) {
            if ($IgnoreDirs -notcontains $_.Name) {
                "$indent$($_.Name)\" | Out-File "./output.txt" -Append
                Show-Tree -path $_.FullName -indent ("$indent`t")
            }
        }
        else {
            if ($IgnoreFiles -notcontains $_.Name) {
                "$indent$($_.Name)" | Out-File "./output.txt" -Append
            }
        }
    }
}

Remove-Item "./output.txt" -ErrorAction SilentlyContinue
Show-Tree "./"
