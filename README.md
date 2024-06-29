# Basic Running

#### configuration

certificates are required to be placed in app directory as cert.pem with their privatekey key.pem
You can generate a self signed key cert pair with the below command
    openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -sha256 -days 3650 -nodes -subj "/C=XX/ST=StateName/L=CityName/O=CompanyName/OU=CompanySectionName/CN=CommonNameOrHostname"

Other configuration is by environment variables below
    BINDADDR  default :3000
    LOGLEVEL  default INFO - can be one of - DEBUG INFO WARN ERROR

#### run latest server
./pullrun.sh
