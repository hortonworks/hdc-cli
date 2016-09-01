# Hortonworks Data Cloud - Command Line Interface

## Usage
To see the available commands `hdc -h`.
```
NAME:
   hdc -

USAGE:
   hdc [global options] command [command options] [arguments...]

VERSION:
   0.1.0-2016-09-01T11:30:37

AUTHOR(S):
   Hortonworks

COMMANDS:
     configure           configure the server address and credentials used to communicate with this server
     create-cluster      creates a new cluster
     list-cluster-types  list the available blueprints
     describe-cluster    list the available clusters
     list-clusters       list the available clusters
     terminate-cluster   terminates a cluster
     help, h             Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --debug        debug mode [$DEBUG]
   --help, -h     show help
   --version, -v  print the version
```
Each command provides a help flag with a description and the accepted flags and subcommands, e.g: `hdc configure -h`.
```
NAME:
   hdc-cli configure - configure the server address and credentials used to communicate with this server

USAGE:
   hdc-cli configure [command options] [arguments...]

DESCRIPTION:
   it will save the provided server address and credential to /Users/khorvath/.hdc/config

OPTIONS:
   --server value    server address [$CB_SERVER_ADDRESS]
   --username value  user name (e-mail address) [$CB_USER_NAME]
   --password value  password [$CB_PASSWORD]

```

### Configure
Although there is an option to provide some global flags to every command to which Cloudbreak to connect to, it is recommended to save the configuration. 
```
hdc configure --server ec2-52-29-224-64...compute.amazonaws.com --username your@email --password your-password
```
This will save the configuration into the user's home directory. To see its content: `cat ~/.hdc/config`. If this config file is present you don't need to specify the connection flags anymore,
otherwise you need to specify these flags to every command.
```
hdc list-clusters --server ec2-52-29-224-64...compute.amazonaws.com --username your@email --password your-password
```

### Create cluster
Most commands have a sub command called `generate-cli-skeleton` and `--cli-input-json` parameter that you can use to store parameters in JSON and read them from a file instead of typing them at the command line.
```
hdc create-cluster generate-cli-skeleton
```
Direct the output to a file to save the skeleton locally.
```
hdc create-cluster generate-cli-skeleton > create_cluster.json
```
To create a cluster fill the empty values or change the existing ones, e.g:
```
{
  "ClusterName": "my-cluster",
  "HDPVersion": "2.5",
  "ClusterType": "EDW-ETL: Apache Spark 2.0-preview",
  "Master": {
    "InstanceType": "m4.xlarge",
    "VolumeType": "gp2",
    "VolumeSize": 32,
    "VolumeCount": 1
  },
  "Worker": {
    "InstanceType": "m3.xlarge",
    "VolumeType": "ephemeral",
    "VolumeSize": 40,
    "VolumeCount": 2
  },
  "InstanceCount": 3,
  "SSHKeyName": "my-existing-keypari-name",
  "RemoteAccess": "0.0.0.0/0",
  "WebAccess": true,
  "ClusterAndAmbariUser": "admin",
  "ClusterAndAmbariPassword": "admin"
}

```
and pass the JSON configuration to the `--cli-input-json` parameter.
```
hdc create-cluster --cli-input-json create_cluster.json
```

### Terminate cluster
To terminate the previously created cluster use the `terminate-cluster` command.
```
hdc terminate-cluster --cluster-name my-cluster
```

### Describe cluster
If you want to clone a cluster the `describe-cluster` command can be useful.
```
hdc-cli describe-cluster --cluster-name my-cluster

{
  "ClusterName": "my-cluster",
  "HDPVersion": "2.5.0.0",
  "ClusterType": "EDW-ETL: Apache Spark 2.0-preview",
  "Master": {
    "InstanceType": "m4.xlarge",
    "VolumeType": "gp2",
    "VolumeSize": 32,
    "VolumeCount": 1
  },
  "Worker": {
    "InstanceType": "m3.xlarge",
    "VolumeType": "ephemeral",
    "VolumeSize": 40,
    "VolumeCount": 2
  },
  "InstanceCount": 3,
  "SSHKeyName": "my-existing-keypari-name",
  "RemoteAccess": "0.0.0.0/0",
  "WebAccess": true,
  "ClusterAndAmbariUser": "admin",
  "ClusterAndAmbariPassword": "admin",
  "Status": "CREATE_IN_PROGRESS",
  "StatusReason": "Creating infrastructure"
}

```

## Debug
To enable the debug logging use the `--debug` global switch
```
hdc --debug list-clusters
```
or provide it as an environment variable `export DEBUG=1` or inline
```
DEBUG=1 hdc list-clusters
```

## Dependency management

This project uses [Godep](https://github.com/tools/godep) for dependency management. You can install it and download the configured dependency versions with:
```
go get github.com/tools/godep
godep restore
```

If you changed the external dependencies then you can save it with:
```
godep save
```