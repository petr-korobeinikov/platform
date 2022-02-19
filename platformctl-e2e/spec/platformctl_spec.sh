Describe 'platformctl'
  It 'available on the system'
    When run command command -v platformctl

    The status should be success
    The output should end with platformctl
  End
End
