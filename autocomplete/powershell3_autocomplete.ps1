$Completion_clc = {
 
    param($commandName, $parameterName, $wordToComplete, $commandAst, $fakeBoundParameter)

    (invoke-expression "$parameterName --generate-bash-completion").Split(" ") |
    ForEach-Object {
        New-Object System.Management.Automation.CompletionResult $_, $_, 'ParameterValue', ('{0} ({1})' -f $_, $_)
    }
}

# Register the handler.
if (-not $global:options) {
    $global:options = @{CustomArgumentCompleters = @{};NativeArgumentCompleters = @{}}
}
$global:options['NativeArgumentCompleters']['clc.exe'] = $Completion_clc

# Override tabexpansion2. 
$function:tabexpansion2 = $function:tabexpansion2 -replace
    'End\r\n{',
    'End { if ($null -ne $options) { $options += $global:options} else {$options = $global:options}'
