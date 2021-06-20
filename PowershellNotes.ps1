#https://github.com/PowerShell/PowerShell/blob/master/docs/learning-powershell/powershell-beginners-guide.md

#ps
Get-Process
Get-Process -Name chrome, code

#gal
Get-Alias cls

#echo
Write-Output $HOME

#cd
Set-Location $HOME
Set-Location $env:GITHUB\learningPowerShell

#dir
Get-ChildItem
Get-ChildItem -Recurse
Get-ChildItem –Path *.txt -Recurse

#touch
New-Item -Path ./test.txt
# force overwrite
New-Item text.txt -Value "asdf" -Force

#write to file
Set-Content -Path ./test.txt -Value "Hello world again!"
"Hello world!!!" > test.txt

#type
Get-Content -Path ./test.txt

#del
Remove-Item ./test.txt

$PSVersionTable

#help
Get-Help -Name Get-Process
Get-Help -Name Get-Process -Examples
Get-Help -Name Get-Process -Full
Get-Help Get-Process -Example

#list available commands
Get-Command
#get syntax
Get-Command Get-Process -Syntax

#sort
Get-ChildItem | Sort-Object -Descending
#largest object
Get-ChildItem | Sort-Object -Property Length -Descending | Select-Object -First 1

#variables
$four = 4
$five = 5
$nine = $four + $five
Write-Output $nine