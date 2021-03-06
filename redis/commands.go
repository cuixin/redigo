package redis

var (
    //---------------- HASH ----------------//
    HGET         = _command("HGET")
    HSET         = _command("HSET")
    HDEL         = _command("HDEL")
    HEXISTS      = _command("HEXISTS")
    HGETALL      = _command("HGETALL")
    HINCRBY      = _command("HINCRBY")
    HINCRBYFLOAT = _command("HINCRBYFLOAT")
    HKEYS        = _command("HKEYS")
    HLEN         = _command("HLEN")
    HMGET        = _command("HMGET")
    HMSET        = _command("HMSET")
    HSETNX       = _command("HSETNX")
    HVALS        = _command("HVALS")
    HSCAN        = _command("HSCAN")
    //---------------- STRING ----------------//
    GET         = _command("GET")
    SET         = _command("SET")
    DEL         = _command("DEL")
    APPEND      = _command("APPEND")
    BITCOUNT    = _command("BITCOUNT")
    BITOP       = _command("BITOP")
    DECR        = _command("DECR")
    GETBIT      = _command("GETBIT")
    GETRANGE    = _command("GETRANGE")
    GETSET      = _command("GETSET")
    INCR        = _command("INCR")
    INCRBY      = _command("INCRBY")
    INCRBYFLOAT = _command("INCRBYFLOAT")
    MGET        = _command("MGET")
    MSET        = _command("MSET")
    MSETNX      = _command("MSETNX")
    PSETEX      = _command("PSETEX")
    SETBIT      = _command("SETBIT")
    SETEX       = _command("SETEX")
    SETNX       = _command("SETNX")
    SETRANGE    = _command("SETRANGE")
    STRLEN      = _command("STRLEN")
    //---------------- SCRIPT ----------------//
    EVAL          = _command("EVAL")
    EVALSHA       = _command("EVALSHA")
    SCRIPT_EXISTS = _command("SCRIPT EXISTS")
    SCRIPT_FLUSH  = _command("SCRIPT FLUSH")
    SCRIPT_KILL   = _command("SCRIPT KILL")
    SCRIPT_LOAD   = _command("SCRIPT LOAD")
    //---------------- CONNECTION ----------------//
    AUTH   = _command("AUTH")
    ECHO   = _command("ECHO")
    PING   = _command("PING")
    QUIT   = _command("QUIT")
    SELECT = _command("SELECT")
    //---------------- SERVER ----------------//
    BGREWRITEAOF     = _command("BGREWRITEAOF")
    BGSAVE           = _command("BGSAVE")
    CLIENT_KILL      = _command("CLIENT KILL")
    CLIENT_LIST      = _command("CLIENT LIST")
    CLIENT_GETNAME   = _command("CLIENT GETNAME")
    CLIENT_SETNAME   = _command("CLIENT SETNAME")
    CONFIG_GET       = _command("CONFIG GET")
    CONFIG_REWRITE   = _command("CONFIG REWRITE")
    CONFIG_SET       = _command("CONFIG SET")
    CONFIG_RESETSTAT = _command("CONFIG RESETSTAT")
    DBSIZE           = _command("DBSIZE")
    DEBUG_OBJECT     = _command("DEBUG OBJECT")
    DEBUG_SEGFAULT   = _command("DEBUG SEGFAULT")
    FLUSHALL         = _command("FLUSHALL")
    FLUSHDB          = _command("FLUSHDB")
    INFO             = _command("INFO")
    LASTSAVE         = _command("LASTSAVE")
    MONITOR          = _command("MONITOR")
    SAVE             = _command("SAVE")
    SHUTDOWN         = _command("SHUTDOWN")
    SLAVEOF          = _command("SLAVEOF")
    SLOWLOG          = _command("SLOWLOG")
    SYNC             = _command("SYNC")
    TIME             = _command("TIME")
    //---------------- SET ----------------//
    SADD        = _command("SADD")
    SCARD       = _command("SCARD")
    SDIFF       = _command("SDIFF")
    SDIFFSTORE  = _command("SDIFFSTORE")
    SINTER      = _command("SINTER")
    SINTERSTORE = _command("SINTERSTORE")
    SISMEMBER   = _command("SISMEMBER")
    SMEMBERS    = _command("SMEMBERS")
    SMOVE       = _command("SMOVE")
    SPOP        = _command("SPOP")
    SRANDMEMBER = _command("SRANDMEMBER")
    SREM        = _command("SREM")
    SUNION      = _command("SUNION")
    SUNIONSTORE = _command("SUNIONSTORE")
    SSCAN       = _command("SSCAN")

    //---------------- LIST ----------------//
    BLPOP      = _command("BLPOP")
    BRPOP      = _command("BRPOP")
    BRPOPLPUSH = _command("BRPOPLPUSH")
    LINDEX     = _command("LINDEX")
    LINSERT    = _command("LINSERT")
    LLEN       = _command("LLEN")
    LPOP       = _command("LPOP")
    LPUSH      = _command("LPUSH")
    LPUSHX     = _command("LPUSHX")
    LRANGE     = _command("LRANGE")
    LREM       = _command("LREM")
    LSET       = _command("LSET")
    LTRIM      = _command("LTRIM")
    RPOP       = _command("RPOP")
    RPOPLPUSH  = _command("RPOPLPUSH")
    RPUSH      = _command("RPUSH")
    RPUSHX     = _command("RPUSHX")

    //---------------- SORTED-SETS ----------------//
    ZADD             = _command("ZADD")
    ZCARD            = _command("ZCARD")
    ZCOUNT           = _command("ZCOUNT")
    ZINCRBY          = _command("ZINCRBY")
    ZINTERSTORE      = _command("ZINTERSTORE")
    ZRANGE           = _command("ZRANGE")
    ZRANGEBYSCORE    = _command("ZRANGEBYSCORE")
    ZRANK            = _command("ZRANK")
    ZREM             = _command("ZREM")
    ZREMRANGEBYRANK  = _command("ZREMRANGEBYRANK")
    ZREMRANGEBYSCORE = _command("ZREMRANGEBYSCORE")
    ZREVRANGE        = _command("ZREVRANGE")
    ZREVRANGEBYSCORE = _command("ZREVRANGEBYSCORE")
    ZREVRANK         = _command("ZREVRANK")
    ZSCORE           = _command("ZSCORE")
    ZUNIONSTORE      = _command("ZUNIONSTORE")
    ZSCAN            = _command("ZSCAN")

    //---------------- TRANSATION ----------------//
    DISCARD = _command("DISCARD")
    EXEC    = _command("EXEC")
    MULTI   = _command("MULTI")
    UNWATCH = _command("UNWATCH")
    WATCH   = _command("WATCH")

    //---------------- PUB/SUB ----------------//
    // PSUBSCRIBE pattern [pattern ...]
    // PUBSUB subcommand [argument [argument ...]]
    // PUBLISH channel message
    // PUNSUBSCRIBE [pattern [pattern ...]]
    // SUBSCRIBE channel [channel ...]
    // UNSUBSCRIBE [channel [channel ...]]

)

func _command(method string) func(conn Conn, args ...interface{}) (interface{}, error) {
    return func(conn Conn, args ...interface{}) (interface{}, error) {
        return conn.Do(method, args...)
    }
}
