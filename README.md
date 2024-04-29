```
      ······························································································
      :                                                                                            :
      :   ____  _____  ___   ____   ___   ____    _____  ____   ___ __     __  ____  _____   ____  :
      :  / _  ||____ ||_ _| / _  | / _ \ |___ \  |____ ||___ \ |_ _|\ \   / / / _  ||____ | |___ \ :
      : | (_| |  |_  | | | | (_| || | | |    | |   |_  |    | | | |  \ \ / / | (_| |  |_  | / ___/ :
      :  > _  | ___| | | |  \__  || |_| | ___| |  ___| | ___| | | |   \ V /   > _  | ___| || (___  :
      : /_/ |_||_____||___|    |_| \___/ |____/  |_____||____/ |___|   \_/   /_/ |_||_____| \____| :
      :                                                                                            :
      ······························································································
```
This helper script is for copying a service from one cluster to another cluster. By running the script, it will download the service file from the source cluster to the local directory which is (~/), and apply it to the destination cluster. This script needs three input parameters to be able to work out. The parameters are source namespace (-n), source service name (-s), destination cluster (-dc), and destination namespace (-dn). Before running the script, you must be at the source cluster where the source service exists. After the successful running of the script, you will be at the destination cluster because the code changes the cluster internally before applying the downloaded service file.
