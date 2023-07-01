package state

// todo use k8 api to get the statefile containing previous resource replica state

// cases
// 1. (resources exist, empty state) -> sleep
//     - create statefile
// 2. (resources exist, statefile exists) -> sleep
// 	   - replace statefile
// 3. (resources exist, empty state) -> wake
//     - create statefile
// 4. (resources exist, statefile exists) -> wake
//     a) current resource replicas > statefile replicas
//        - update statefile to current resource replicas
//		  - delete statefile
//     b) current resource replicas <= statefile replicas
//        - do nothing
//		  - delete statefile

// 5. (resources dont exist, empty state) -> sleep
//     - do nothing
// 6. (resources dont exist, statefile exists) -> sleep
//     - do nothing
// 7. (resources dont exist, empty state) -> wake
//     - do nothing
// 8. (resources dont exist, statefile exists) -> wake
//     - create resources using statefile
//	   - delete statefile
