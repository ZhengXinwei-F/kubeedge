## Keadm initializes multiple instances of cloudcore

- [Summary](#summary)
- [Motivation](#motivation)
  - [Goals](#goals)
  - [Non-goals](#non-goals)
- [Design Details](#design-details)
  - [Architecture](#architecture)
- [Test Cases](#test-cases)
- [Plan](#plan)

### Summary

目前，keadm只支持安装单实例的cloudcore，其中

### Motivation

To meet the user's installation requirements for more cni, the edge node should be able to acquire kubernetes version information.

The cni plugin will be able to get cluster version information through the MetaServer with this feature. Further development can be carried out if the user requires other non-resource uri response.

##### Goals

- MetaServer allows for the transparent transmission of non-resource URIs.  And save the results in the sqlite database.
- If the connection between the edge node and the cluster is lost, MetaServer can get the cache from sqlite to avoid an error return.

##### Non-goals

- Transmit uris other than /version in MetaServer.

### Design Details

#### Architecture

![](../images/edge-version/edge-version.png)

##### On the cloud：

- Verify non-resource requests using the uri+verb whitelist.
- Non-resource requests are routed through restClient without regard for the verb of the request.


##### On the edge

- Verify non-resource requests using the uri+verb whitelist.
- Regardless of the selected verb, the non-resource request is sent a message to cloudhub with uri as the key.
- The key for storage and querying in the local database is uri.

##### Before using

- Enable `dynamicController`

Dynamic modules is set to false by default; modify it to true as shown below.

cloudcore.yaml:
```
      dynamicController: 
        enable: true      
```

Enable `MetaServer`

edgecore.yaml:
```
  metaManager:
    metaServer:
      enable: true
```

### Test Cases

1. The application on the edge node connects to the /version interface via MetaServer's port.
2. Disconnect the network connection between cloud and edge，and try the /version request.

### Plan



