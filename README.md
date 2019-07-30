# _Google Calender_ OMG Microservice

[![Open Microservice Guide](https://img.shields.io/badge/OMG%20Enabled-👍-green.svg?)](https://microservice.guide)


An OMG service for google calender, it allows users to create and edit events. Reminders can be enabled for events, with options available for type and time. Event locations can also be added, and other users can be invited to events.

## Direct usage in [Storyscript](https://storyscript.io/):

##### Event
```coffee
amplitude event userId:'abc@example.com' eventType:'event name' eventProperties:'{"property1":"one","property2":"two"}' userProperties:'{"Cohort":"Test A"}'appVersion:'1.0.0' platform:'ios' osName:'ios' osVersion:'0.13' deviceBrand:'iphone' deviceManufacturer:'Apple' deviceModel:'IPhone X' city:'Pune' country:'India' region:'India' dma:'India' language:'English' price:1590.5 quantity:1 revenue:1590.52 productId:'P123' revenueType:'refund' ip:'127.0.0.1'
{"success": "true","message": "Event uploaded successfully","statusCode": 200}
```

Curious to [learn more](https://docs.storyscript.io/)?

✨🍰✨

## Usage with [OMG CLI](https://www.npmjs.com/package/omg)

##### Event
```shell
$ omg run event -a userId=<USER_ID> -a eventType=<EVENT_TYPE> -a eventProperties=<MAP_TYPE_EVENT_PROPERTIES> -a userProperties=<MAP_TYPE_USER_PROPERTIES> -a appVersion=<APP_VERSION> -a platform=<PLATFORM> -a osName=<OS_NAME> -a osVersion=<OS_VERSION> -a deviceBrand=<DEVICE_BRAND> -a deviceManufacturer=<DEVICE_MANUFACTURER> -a deviceModel=<DEVICE_MODEL> -a city=<USER_CITY> -a country=<USER_COUNTRY> -a region=<USER_REGION> -a dma=<DESIGNATED_MARKET_AREA> -a language=<USER_SELECTED_LANGUAGE> -a price=<ITEM_PRICE> -a quantity=<QUANTITY> -a revenue=<REVENUE> -a productId=<PRODUCT_ID> 
-a revenueType=<REVENUE_TYPE> -a ip=<IP_ADDRESS> -e API_KEY=<API_KEY>
```

**Note**: the OMG CLI requires [Docker](https://docs.docker.com/install/) to be installed.

## License
[MIT License](https://github.com/omg-services/amplitude/blob/master/LICENSE).
