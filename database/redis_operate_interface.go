package database

/*this file contains the interface of redis operation,
* you can use functions below to interact with redis
 */
import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var cxt = context.Background()

/*function: connect to the default redis,
* params: do not need in-params,
* return: a pointer of redis.client while the redis is the default one
 */
func ConnectToRedis() *redis.Client {
	var opt = redis.Options{
		Addr:     "redis.DianasDog.secoder.local:6379",
		Password: "",
		DB:       0}
	var client = redis.NewClient(&opt)
	return client
}

/*function: insert string type key-value to redis, if key is existed, the value will be updated.
* params: three params are all necessary, the first type is '*redis.Client' which is 'a client of your target redis',
* the second type is 'string' which is 'the key of your data', the third type is 'string' which is 'the value of your data'
* return: a bool value. true means the insertion is successful; false means error generates in the insertion process no
* matter the error will shut the program or not.
 */
func SetToRedis(client *redis.Client, key string, value string) bool {
	err := client.Set(cxt, key, value, 0).Err()
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

/*function: get value from redis with needing-provided key
* params: two params are all necessary, the first type is '*redis.Client' which is 'a client of your target redis',
* the second type is 'string' which is 'the key you provide'.
* return: two params. the first is the value corresponding to your key. the second is a bool value, true means the get
* operation is successful, false means error generates in the get process.
 */
func GetFromRedis(client *redis.Client, key string) (string, bool) {
	str, err := client.Get(cxt, key).Result()
	if err != nil {
		fmt.Println(err)
		return str, false
	}
	return str, true
}

/*function: delete key-value element with the key you provide
* params: two params are both necessary, the first type is '*redis.Client' which is 'a client of your target redis',
* the second type is 'string' which is 'the key you provide'.
* return: a bool value. true means the deletion is successful; false means error generates in the deletion process no
* matter the error will shut the program or not.
 */
func DeleteFromRedis(client *redis.Client, key string) bool {
	err := client.Del(cxt, key).Err()
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

/*function: judge a key whether exist in the redis
* params: two params are both necessary, the first type is '*redis.Client' which is 'a client of your target redis',
* the second type is 'string' which is 'the key you provide'.
* return: two bool values. the first represents the key's existence in the redis, the second represents this process
* is successful or not.
 */
func ExistInRedis(client *redis.Client, key string) (bool, bool) {
	exist, err := client.Do(cxt, "EXISTS", key).Bool()
	if err != nil {
		fmt.Println(err)
		return exist, false
	}
	return exist, true
}
