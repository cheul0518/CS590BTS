#### Seungchuel Kim

#### Professor Kate

#### CS590BTS

#### March 1, 2019

<br />

### Task 1

- Task 1.1 & 2 Generate Network Artifacts and Edit Docker Compose Files

    - This network is modifed from the first-network for three organizations, each with one peer node, and a solo ordering service. Matching the required structure, artifacts, and genesis block, anchor peer transition for each org, and channel creation are successfully generated and printed. The first network totally works

<img src = "Images/task11.png">

<img src = "Images/Task122.png">

<img src = "Images/task121.png">

- Task 1.3 Run Sample Chaincodes

    - The chaincode succesfully runs as well as querying and invoking.

<img src = "Images/task131.png">

<img src = "Images/task132.png">

<br />

### Task 2


- Task 2.1 & 2.2

/*
Thanks a lot for the link, and for the task 2 my understanding is: 1) understand the syntax for Golang and check the samples in the “balance-transfer” directory, they gave us sample steps. We can replace that we our chaincodes to debug; 2) Write the supply chain management chain codes for car supply chain: we have [supplier] [manufacturere] [dealer] and [car], 4 different roles, and each of them are granted with different operations (chain codes in different go scripts). There are 7 functions with different business logics, which are the chain codes we need to write, test, and install onto our network.

I think components are just parts/pieces in a car, and all these 4 role types can access the chaincode with different endorsement policies


Say we have tires and airbags as our components, then 1) first, for example, supplier0 will call `AddComponent(supplier0, tire0)’ and `AddComponent(supplier0, tire1)`; 2) `TransferComponent(supplier0, manuf0, tire0)` and `TransferComponent(supplier0, manuf0, tire1)` will be called to transfer these two tires to the manufactory; 3) the manufactory will then call MountComponent(manuf0, tire0, car0) and MountComponent(manuf0, tire1, car0) and something like that…

So the basic idea will be:
Supplier (SP) -> Build the component [AddCompoment] -> Transfer it to manufacture (MF) [TransferComponent] -> (MF) mount these components to cars [MountComponent] OR (MF) transfer these components to car dealers [TransferComponent] OR (MF) can call [RecallComponent]  if it is not retired, i.e. it is not replaced before; Authorized dealer (AD) can call [TransferComponent] OR [ReplaceComponent]; only cars (C) can call [CheckComponent]

So the basic idea will be:
Supplier (SP) ->
 Build the component [AddCompoment] ->
 Transfer it to manufacture (MF) [TransferComponent] ->
 (MF) mount these components to cars [MountComponent] OR (MF) transfer these components to car dealers [TransferComponent] OR (MF) can call [RecallComponent]  if it is not retired, i.e. it is not replaced before -> 
Authorized dealer (AD) can call [TransferComponent] OR [ReplaceComponent] ->
only cars (C) can call [CheckComponent]
*/


<br />

### Task 3
