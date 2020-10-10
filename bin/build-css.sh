#!/bin/bash

## Uses `scss` to compile the stylesheets in `web/stylesheets`
## Requires SCSS available on the path

set -euo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd $dir/..

sass --load-path=../npn/npnasset/stylesheets --no-source-map web/stylesheets/style.scss web/assets/vendor/rituals.css
sass --load-path=../npn/npnasset/stylesheets --style=compressed --no-source-map web/stylesheets/style.scss web/assets/vendor/rituals.min.css
