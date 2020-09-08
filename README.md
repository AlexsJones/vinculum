# vinculum

Vinculum is a command and control system for low powered devices that cannot run k3s or docker swarm. 
It uses a single binary in a leader/follower mode with a grpc protocol to send commands.
There is a small amount of intelligence in scheduling and should be thought of as a shortcut from ansible/ssh commands.

As an example this works well across a cluster of Raspberry Pi Zero compute units.

## Install
```shell script
 curl https://raw.githubusercontent.com/AlexsJones/vinculum/master/install.sh | sh -
```

### Run

Start the leader

```
vinculum leader listen
```

Start a follower

```
vinculum follower connect -s localhost:7559
```


