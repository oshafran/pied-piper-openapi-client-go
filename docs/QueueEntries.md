# QueueEntries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entries** | Pointer to [**[]OnDemandQueueEntry**](OnDemandQueueEntry.md) |  | [optional] 

## Methods

### NewQueueEntries

`func NewQueueEntries() *QueueEntries`

NewQueueEntries instantiates a new QueueEntries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueueEntriesWithDefaults

`func NewQueueEntriesWithDefaults() *QueueEntries`

NewQueueEntriesWithDefaults instantiates a new QueueEntries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntries

`func (o *QueueEntries) GetEntries() []OnDemandQueueEntry`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *QueueEntries) GetEntriesOk() (*[]OnDemandQueueEntry, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *QueueEntries) SetEntries(v []OnDemandQueueEntry)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *QueueEntries) HasEntries() bool`

HasEntries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


