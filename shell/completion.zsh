#compdef xraybuilder

function _xraybuilder {
  _arguments \
    '-c[Specify xray config path]:config file:_files' \
    '-k[Specify xray keypair path]:keypair file:_files' \
    '-v[Enable verbose mode]' \
    '1:subcommand:(user setup install-misc)' \
    '*::arg:->args'

  case $state in
    args)
      case $line[1] in
        user)
          _xraybuilder_user
          return 0
          ;;
        setup)
          _xraybuilder_setup
          return 0
          ;;
        install-misc)
          _nothing
          return 0
          ;;
      esac
      ;;
    user_args)
      case $line[2] in
        add)
          _xraybuilder_user_add
          return 0
          ;;
        remove)
          _xraybuilder_user_remove
          return 0
          ;;
        share)
          _xraybuilder_user_share
          return 0
          ;;
        list)
          _nothing
          return 0
          ;;
      esac
      ;;
  esac

}

function _xraybuilder_user {
  _arguments \
    '1:subcommand:(add remove share list)' \
    '*::arg:->user_args'
}

function _xraybuilder_setup {
  _arguments \
    '-d[Specify destination]:destination URL:_urls' \
    '-i[Install xray-core by version]:version number:'
}

function _xraybuilder_user_add {
  _arguments '1:comment:Comment of the new user'
}

function _xraybuilder_user_remove {
  _arguments '1:id_or_comment:Id or Comment of the user to be removed'
}

function _xraybuilder_user_share {
  _arguments \
    '1:id_or_comment:Id or Comment of the user to share' \
    '(-f --format)'{-f,--format}'[Specify format for sharing]:format:(qr json link)' \
    '*:filename:_files'
}

compdef _xraybuilder xraybuilder
