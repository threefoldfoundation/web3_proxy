# GridProxy API client
Easily access Threefold grid APIs from vlang. gridproxy is v module include the API client along with API-specific information such as the root URL for the different networks available in the threefold grid. They also include classes that represent entities in the context of the API in sub-module `model`, and that are useful for making conversions between JSON objects and V objects. and some types with helper methods to convert the machine-friendly units returned by the API to more human-friendly units.

## client usage

If you want to use the client, you need to import it in your code.

* import the client:
  ```v
  import threefoldtech.vgrid.gridproxy
  ```

* create a client:
  ```v
  // create a client for the testnet, with API cache disabled
  // you can pass true as second arg to enable cache
  mut gp_client := gridproxy.get(.test, false)!
  ```

* use the client to interact with the gridproxy API:
  ```v
  // get farm list
  farms := gp_client.get_farms()! // you should handle any possible errors in your code
  // get gateway list
  gateways := gp_client.get_gateways()!
  // get node list
  nodes := gp_client.get_nodes()!
  // get contract list
  contracts := gp_client.get_contracts()!
  // get grid stats
  stats := gp_client.get_stats()!
  // get node by id
  node := gp_client.get_node_by_id(u64(16))!
  // get node stats
  node_stats := gp_client.get_node_stats_by_id(u64(16))!
  // get twins
  twins := gp_client.get_twins()!
  ```
  for all available methods on the client, see [GridProxy API client modules doc][GridProxy API client modules doc](../../../3bot/vgrid/gridproxy/docs/)

* filtering:
  ```v
  // getting only dedicated farms
  farms_dedicated := gp_client.get_farms(dedicated: true)!
  // getting only farms with at least one free ip
  farms_with_free_ips := gp_client.get_farms(free_ips: u64(1))!
  // pagination options:
  // get first page of farms
  farms_first_page := gp_client.get_farms(page: u64(1))!
  // you can mix any filters and pagination options
  farms_first_page_dedicated := gp_client.get_farms(page: u64(1), dedicated: true)!
  // access the field of first farm in the list
  // the API could return an empty list if no farm is found
  // you should handle this case in your code
  if farms_first_page.len > 0 {
    println(farms_first_page[0].name)
  }
  ```

  for all available filters, see [GridProxy API client modules doc][GridProxy API client modules doc](../../../3bot/vgrid/gridproxy/docs/)

* helper methods:
  ```v
  node := nodes[0]
  node.updated_at // 1655940222
  node.created // 1634637306
  // you can convert the timestamp to V Time object easily with the helper method
  node.created.to_time() // 2021-10-19 09:55:06
  node.created.to_time().local() // 2021-10-19 11:55:06
  node.created.to_time().relative() // last Oct 19
  node.created.to_time().relative_short() // 246d ago
  // lets check another field with different type
  node.uptime // 18958736
  // you can convert the seconds to a human-readable duration with the helper method
  node.uptime.to_days() // 219.42981481481482
  node.uptime.to_hours() // 5266.315555555556
  node.uptime.to_minutes() // 315978.93333333335
  // now to the capacity helper methods
  node.total_resources.mru // 202803036160
  // you can `to_megabytes`, `to_gigabytes` and `to_terabytes` methods on any resources field.
  node.total_resources.mru.to_gigabytes() // 202.80303616
  // the helper methods available for the billing to help you convert the TFT units as well
  ```
  for all available helper methods, see [GridProxy API client modules doc](../../../3bot/vgrid/gridproxy/docs/)

  TODO:
  * Documented the client iterators and higher-level methods

## Client Examples
there are scripts available to serve as examples in the [examples](../../../3bot/vgrid/examples/) directory. [Docs](../../../3bot/vgrid/examples/)
