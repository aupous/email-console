# Email Console

## Run project

### Run with code

```bash
# Run this command on project directory
go run cmd/main.go email_template.json customers.csv emails/ errors.csv
```

### Run with docker image

```bash
# Pull docker image
docker pull aupous/email-console

# Run bash in the downloaded image
docker run --rm -it --entrypoint /bin/sh aupous/email-console

# Run application
./app email_template.json customers.csv emails/ errors.csv
```
