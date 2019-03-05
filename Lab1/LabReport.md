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

        - The goal is to write my own chaincodes for supplier, manufacturer, dealer, and car. There are six functions: AddComponent(), TransferComponent(), MountComponent(), ReplaceComponent(), RecallComponent(), CheckComponent(). And they are supposed to be accessible under different endorsement policies. 
        - Here is how they work. Supplier builds a component, and transfer it to a manufacturer. The manufacturer either mounts this component to a car or transfers it to a dealer. There would be as well a chance the manucaturer will replace the component with another one. The car dealer either transfers it to a car or replace it with another one. The car checks itself only.





<br />

### Task 3
