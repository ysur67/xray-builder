#compdef xray-builder

function _xray-builder {
  _arguments \
    '-c[Specify xray config path]:config file:_files' \
    '-k[Specify xray keypair path]:keypair file:_files' \
    '-v[Enable verbose mode]' \
    '1:subcommand:(user setup)' \
    '*::arg:->args'

  case $state in
    args)
      case $line[1] in
        user)
          _xray-builder_user
          return 0
          ;;
        setup)
          _xray-builder_setup
          return 0
          ;;
      esac
      ;;
    user_args)
      case $line[2] in
        add)
          _xray-builder_user_add
          return 0
          ;;
        remove)
          _xray-builder_user_remove
          return 0
          ;;
        share)
          _xray-builder_user_share
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

function _xray-builder_user {
  _arguments \
    '1:subcommand:(add remove share list)' \
    '*::arg:->user_args'
}

function _xray-builder_setup {
  _arguments \
    '-d[Specify destination]:destination URL:_urls' \
    '-i[Install xray-core by version]:version number:'
}

function _xray-builder_user_add {
  _arguments '1:comment:Comment of the new user'
}

function _xray-builder_user_remove {
  _arguments '1:id_or_comment:Id or Comment of the user to be removed'
}

function _xray-builder_user_share {
  _arguments \
    '1:id_or_comment:Id or Comment of the user to share' \
    '(-f --format)'{-f,--format}'[Specify format for sharing]:format:(qr json link)' \
    '*:filename:_files'
}

compdef _xray-builder xray-builder
