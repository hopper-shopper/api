./scripts/build.sh

TAR_FILE=caprover.tar

tar czf $TAR_FILE captain-definition bin/
export $(grep -v '^#' staging.env | xargs) && caprover deploy --tarFile $TAR_FILE -a $CAPROVER_APP -u $CAPROVER_URL --appToken $CAPROVER_TOKEN
rm $TAR_FILE
