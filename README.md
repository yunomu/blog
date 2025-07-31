# blog

## Diagram

```mermaid
graph LR
  subgraph blog
    Article[(ArticleTable)]
    User[(UserTable)]
    File[(FileTable)]
    Storage[(PublishBucket)]
    API[[Api]]
    ArticleFunction
    UserFunction
    FileFunction
    ImageReduceFunction
    AuthClient
    GenerateFunction
    CDN[[Distribution]]
    ConsoleCDN[[ConsoleDistribution]]
    ConsoleBucket[(ConsoleBucket)]
  end

  Owner((Owner))
  Internet((Internet))

  Owner -- "read/post/update/publish" --> ConsoleCDN
  ConsoleCDN -- "post/update/publish" --> API
  ConsoleCDN -- "read" --> ConsoleBucket
  ArticleFunction --> Article
  UserFunction --> User
  API --> ArticleFunction
  API --> UserFunction
  API --upload--> FileFunction
  API -.- AuthClient
  AuthClient ---> IDProvider
  API -- publish --> GenerateFunction
  Article --> GenerateFunction
  User --> GenerateFunction
  FileFunction --> Storage
  GenerateFunction --> Storage
  FileFunction --> File
  Storage -.- CDN
  File -.put event.-> ImageReduceFunction
  Storage --> ImageReduceFunction
  ImageReduceFunction --> Storage
  CDN -- view--- Internet
```
## Table: User
### Schema
|AttrName|Type|Key  |
|--------|----|-----|
|Id      |S   |Hash |
|Attr    |S   |Range|
|Name    |S   |     |
### Attributes
* **Id**: User ID
* **Attr**: Always `Main`
* **Name**: Unique user name to be used in URL path
### Index: NameIndex
|AttrName|Key  |
|--------|-----|
|Name    |Hash |
## Table: Article
TODO
## Table: File
### Schema
|AttrName|Type|Key  |
|--------|----|-----|
|Key     |S   |Hash |
|Attr    |S   |Range|
|UserId  |S   |     |
|Name    |S   |     |
|CType   |S   |     |
|TS      |N   |     |
|Size    |N   |     |
|W       |N   |     |
|H       |N   |     |
|Status  |S   |     |
|UserIdx |S   |     |
### Attributes
* **Key**: S3 object key. Not reused. randomized string.
* **Attr**: Attributes of this object
  * **Orig**: Original file
  * **Rep:{W}x{H}**: Reduced replica
* **UserId**: User ID by IDP
* **UserIdx**: Same as `UserId` when `Attr` is `Orig`. Otherwise, empty.
* **Name**: Filename
* **CType**: Content Type
* **TS**: Created at
* **Size**: Image size(bytes)
* **W**,**H**: Widht and height
* **Status**: reserved|available|deleting|deleted
### Index: UserIndex
|AttrName|Key  |
|--------|-----|
|UserIdx |Hash |
|TS      |Range|

## URL path design
* **/{User:Id}**: Index
* **/{User:Id}/{Article:Id}**: Article
* **/files/{File:Key}/{File:Attr}.{File:Ext}**
