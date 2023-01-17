echo "Building elookupbe"
cd e-lookup-be/web
go build -o elookup

cd ../../

echo "Starting e-lookup-be"
ZINC_SEARCH_SERVER_URL="http://localhost:4080/api/" ZINC_SEARCH_USER="admin" ZINC_SEARCH_PASSWORD="Complexpass#123" ZINC_SEARCH_INDEX_NAME="enron-index" ./e-lookup-be/web/elookup &

echo "Installing dependencies for e-lookup-fe"
cd e-lookup-fe
npm i

echo "Starting e-lookup-fe"
VITE_ELOOKUP_BACKEND_QUERY_URL="http://localhost:3000/api/v1/lookup?" npm run dev &

cd ..

echo "Starting ZincSearch server"
ZINC_FIRST_ADMIN_USER=admin ZINC_FIRST_ADMIN_PASSWORD=Complexpass#123 ./zinc/zinc &
