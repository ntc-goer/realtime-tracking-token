### How to run source
In folder /dev . Run all service by following command \
`docker-compose up`

### Working Time Total
~ 16h including research about blockchain, meta marks , address , transaction , v..v..
### API
There are some exported api by api service
1. GET: http://localhost:8080/subscribes \
=>  Get monitoring address list
> Example: curl --location 'localhost:8080/subscribes'
2. POST: http://localhost:8080/subscribes  \
=> Subscribe an address
> Example: curl --location 'localhost:8080/subscribe' \
   --header 'Content-Type: application/json' \
   --data '{
   "address": "0x535f548601FEff5586388E620fFe280259eC8f0D"
   }'
3. POST: http://localhost:8080/unsubscribe \
=> UnSubscribe an address
> Example: curl --location 'localhost:8080/unsubscribe' \
   --header 'Content-Type: application/json' \
   --data '{
   "address": "0x535f548601FEff5586388E620fFe280259eC8f0D"
   }'
4. GET: http://localhost:8080/transaction \
=> Get all transaction related to address since subscribe time
> Example: curl --location 'localhost:8080/transactions?address=0x90f4b3Fac242082662DD4f80793141fb35e4FA6d'

