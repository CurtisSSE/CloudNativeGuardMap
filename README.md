## University of Bath 2025 - Dissertation Project Prototype
### Cloud Native Guard Map
#### The automated threat modeling toolkit.

#### Disclaimer:

This is a very rough prototype aimed at demonstrating how threat modelling can be done in a more automated manner. Please do not make pull requests/comments about authentication lifecycles and the way global states are written. For this demonstration, it is intentional as the entire application is client-side and used for demonstration purposes only. This is not a commercial or production level application.

#### Demo video can be found here: https://1drv.ms/v/c/4a764f8c7387a09c/Ea92E_I7XExHl4O_8uCPAfkBCmGs16JnaFu9Ne2JSTwgJg?e=9TZuEM

Manual threat modeling takes time. It requires you to look over all of the assets you want to model.

This application changes that and does the heavy lifting for you, as well as the lifting of the brush to a canvas to draw a perfectly aligned threat model.

**This is a dissertation project prototype. It works under the scenarios detailed in the 'Results' section of the dissertation and can only currently handle a single subscription, single nested resource group and as many VMs and NSGs fit on the canvas page.**

##### To run this project:

**Backend**: From within the /CloudNativeGuardMap repository, run **go run main.go**

**Frontend**: From within the /CloudNativeGuardMap/Site repository, run **npm run dev**

**DO NOT RUN USING PRODUCTION CREDENTIALS, OR AZURE CREDENTIALS FOR ANY LIVE ENVIRONMENT WITH PII OR SENSITIVE DATA OF ANY KIND. THIS PROJECT IS ONLY SUITABLE FOR PURPOSE-BUILT DEVELOPMENT ENVIRONMENTS.**

