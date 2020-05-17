script_directory=$(dirname ${BASH_SOURCE[0]})
src_directory=$(dirname "$script_directory")

source $script_directory/variables.sh

while test $# -gt 0; do
  case "$1" in
  -t) ;&
  --target)
    case "$2" in
    all)
      target="all"
      ;;
    api)
      target="api"
      ;;
    web)
      target="web"
      ;;
    *)
      echo "Invalid target: $2"
      exit 0
      ;;
    esac
    shift
    ;;
  -e) ;&
  --env) ;&
  --environment)
    case "$2" in
    dev) ;&

    development)
      env="development"
      ;;
    prod) ;&
    production)
      env="production"
      ;;
    *)
      echo "Invalid environment: $2"
      exit 0
      ;;
    esac
    shift
    ;;
  -h) ;&
  --help)
    echo "Usage:"
    echo "    ./build.sh [options]"
    echo
    echo "Options:"
    echo "    -t, --target:"
    echo "        all"
    echo "        api"
    echo "        web"
    echo "    -e, --env, --environment:"
    echo "        development (or dev)"
    echo "        production  (or prod)"
    echo

    exit 0
    ;;
  *)
    echo "argument $1"
    ;;
  esac
  shift
done
