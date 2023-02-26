# Golang framework agnostic boilerplate for web api

My motivation for creating this project is to be able to easily switch between different http frameworks without having to rewrite a lot of code, and even allowing the option to not use any framework at all. The limitation of this boilerplate is that code replacement still needs to be done in parts that have a close dependency with the used framework, such as in routing and middleware. I impose this limitation because I still want the ease of using the provided framework parts without having to rewrite or create adapters.

The architecture that I used as a reference in this boilerplate is Clean Architecture. The part that I want to preserve when changing frameworks is the controller to the entity. I use an interface for the presenter part and implement an adapter for each framework.

- [ ] http handler
  - [ ] independent controller
  - [ ] echolabstack adapter
  - [ ] fiber adapter

