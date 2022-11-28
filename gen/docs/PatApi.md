# PatApi

All URIs are relative to *http://localhost:8080/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**createPat**](PatApi.md#createPat) | **POST** /Pats | Метод добавления нового питомца в каталог
[**deletePatById**](PatApi.md#deletePatById) | **DELETE** /pats/{id} | Метод удаления игрушки по идентификатору
[**getPatByStatus**](PatApi.md#getPatByStatus) | **GET** /pats/status/{status} | Метод получения питомца по состоянию
[**getPats**](PatApi.md#getPats) | **GET** /Pats | Метод получения питомца
[**getPayById**](PatApi.md#getPayById) | **GET** /pats/{id} | Метод получения питомца по идентификатору
[**updatePat**](PatApi.md#updatePat) | **PUT** /pats/{id} | Метод обновления питомца в каталоге


<a name="createPat"></a>
# **createPat**
> Pat createPat(pat)

Метод добавления нового питомца в каталог

Метод предназначен для сохранения в БД данных по новому питомцу.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.models.*;
import org.openapitools.client.api.PatApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://localhost:8080/api/v1");

    PatApi apiInstance = new PatApi(defaultClient);
    Pat pat = new Pat(); // Pat | 
    try {
      Pat result = apiInstance.createPat(pat);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling PatApi#createPat");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pat** | [**Pat**](Pat.md)|  |

### Return type

[**Pat**](Pat.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Подтверждение успешного сохранения |  -  |
**400** | Некорректные входные данные. Возвращает список атрибутов с ошибками |  -  |
**0** | Любая неожиданная ошибка |  -  |

<a name="deletePatById"></a>
# **deletePatById**
> deletePatById(id)

Метод удаления игрушки по идентификатору

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.models.*;
import org.openapitools.client.api.PatApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://localhost:8080/api/v1");

    PatApi apiInstance = new PatApi(defaultClient);
    String id = "id_example"; // String | Идентификатор питомца
    try {
      apiInstance.deletePatById(id);
    } catch (ApiException e) {
      System.err.println("Exception when calling PatApi#deletePatById");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **String**| Идентификатор питомца |

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Успешное удаление |  -  |
**0** | Любая неожиданная ошибка |  -  |

<a name="getPatByStatus"></a>
# **getPatByStatus**
> Pat getPatByStatus(status)

Метод получения питомца по состоянию

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.models.*;
import org.openapitools.client.api.PatApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://localhost:8080/api/v1");

    PatApi apiInstance = new PatApi(defaultClient);
    String status = "status_example"; // String | Статус (состояние) питомца
    try {
      Pat result = apiInstance.getPatByStatus(status);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling PatApi#getPatByStatus");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **String**| Статус (состояние) питомца | [enum: продажа, на реализации, возврат]

### Return type

[**Pat**](Pat.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Питомец |  -  |
**0** | Любая неожиданная ошибка |  -  |

<a name="getPats"></a>
# **getPats**
> List&lt;Pat&gt; getPats()

Метод получения питомца

Метод предназначен для получения списка всех питомцев, сохраненных в БД.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.models.*;
import org.openapitools.client.api.PatApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://localhost:8080/api/v1");

    PatApi apiInstance = new PatApi(defaultClient);
    try {
      List<Pat> result = apiInstance.getPats();
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling PatApi#getPats");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**List&lt;Pat&gt;**](Pat.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Список питомцев |  -  |
**0** | Любая неожиданная ошибка |  -  |

<a name="getPayById"></a>
# **getPayById**
> Pat getPayById(id)

Метод получения питомца по идентификатору

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.models.*;
import org.openapitools.client.api.PatApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://localhost:8080/api/v1");

    PatApi apiInstance = new PatApi(defaultClient);
    String id = "id_example"; // String | Идентификатор питомца
    try {
      Pat result = apiInstance.getPayById(id);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling PatApi#getPayById");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **String**| Идентификатор питомца |

### Return type

[**Pat**](Pat.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Питомец |  -  |
**0** | Любая неожиданная ошибка |  -  |

<a name="updatePat"></a>
# **updatePat**
> Pat updatePat(id, pat)

Метод обновления питомца в каталоге

Метод предназначен для обновления в БД данных по имеющейся питомце.

### Example
```java
// Import classes:
import org.openapitools.client.ApiClient;
import org.openapitools.client.ApiException;
import org.openapitools.client.Configuration;
import org.openapitools.client.models.*;
import org.openapitools.client.api.PatApi;

public class Example {
  public static void main(String[] args) {
    ApiClient defaultClient = Configuration.getDefaultApiClient();
    defaultClient.setBasePath("http://localhost:8080/api/v1");

    PatApi apiInstance = new PatApi(defaultClient);
    String id = "id_example"; // String | Идентификатор питомца
    Pat pat = new Pat(); // Pat | 
    try {
      Pat result = apiInstance.updatePat(id, pat);
      System.out.println(result);
    } catch (ApiException e) {
      System.err.println("Exception when calling PatApi#updatePat");
      System.err.println("Status code: " + e.getCode());
      System.err.println("Reason: " + e.getResponseBody());
      System.err.println("Response headers: " + e.getResponseHeaders());
      e.printStackTrace();
    }
  }
}
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **String**| Идентификатор питомца |
 **pat** | [**Pat**](Pat.md)|  |

### Return type

[**Pat**](Pat.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Подтверждение успешного обновления |  -  |
**400** | Некорректные входные данные. Возвращает список атрибутов с ошибками |  -  |
**0** | Любая неожиданная ошибка |  -  |

