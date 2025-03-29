## University of Bath 2025 - Dissertation Project Prototype
### Cloud Native Guard Map
#### The automated threat modeling toolkit.

Manual threat modeling takes time. It requires you to look over all of the assets you want to model.

This application changes that and does the heavy lifting for you, as well as the lifting of the brush to a canvas to draw a perfectly aligned threat model.

**This is a dissertation project prototype. It works under the scenarios detailed in the 'Results' section of the dissertation and can only currently handle a single subscription, single nested resource group and as many VMs and NSGs fit on the canvas page. No AI generated code was used in the creation of this project. All code is original and proprietary unless otherwise stated.**

##### To run this project:

**Backend**: From within the /CloudNativeGuardMap repository, run **go run main.go**

**Frontend**: From within the /CloudNativeGuardMap/Site repository, run **npm run dev**

**DO NOT RUN USING PRODUCTION CREDENTIALS, OR AZURE CREDENTIALS FOR ANY LIVE ENVIRONMENT WITH PII OR SENSITIVE DATA OF ANY KIND. THIS PROJECT IS ONLY SUITABLE FOR PURPOSE-BUILT DEVELOPMENT ENVIRONMENTS.**

