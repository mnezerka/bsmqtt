# BSMQTT Client

Simple MQTT client built as single static binary (no library dependences)

## Configuration

There are three alternatives for configuration of client:

1. Command line flags:

   ```bash
   bsmqtt --mqtt-url ssl://example.net:8883  --mqtt-user bob --mqtt-password secret-password
   ```

1. Configuration file in yaml format (default name and location is `bsmqtt.yaml`
   in your home directory), format is following:

   ```yaml
   mqtt.url: "ssl://example.net:8883"
   mqtt.user: bob
   mqtt.password: "secret-password"
   log.level: DEBUG
   ```
   
   Tool will find this file automatically. If you want to read configuration
   from non-standard location, specify it as:

   ```bash
   bsmqtt --config path-to-yaml-file
   ```

1. Environment variables - each variable has prefix `BSMQTT_` followed by a
   parameter name in capitals. All dots and dashes are replaced by underscores: 

   ```bash
   export BSMQTT_MQTT_URL=ssl://example.net:8883
   export BSMQTT_MQTT_USER=bob
   export BSMQTT_MQTT_PASSWORD=secret-password
   export BSMQTT_LOG_LEVEL=DEBUG
   ```

## Use

Subscribe to `some-topic`:

```bash
bsmqtt sub some-topic
```

Publish `some-message` to `some-topic`:

```bash
bsmqtt  pub some-topic some-message
```
