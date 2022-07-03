# Summary - Enter a brief one-line description of the bug

## Details

This section lays out the details of the bug. Begin with a description
of the buggy behavior.

> The `/metrics` endpoint produces an address location when
> the `bysize` array is displayed in the `/metrics` endpoint `0X0583`

What is the expected behavior?

> The `bysize` array in the `/metrics` endpoint is supposed to display an array
> of objects when rendered as a JSON string in the HTTP response

What steps did you take to reproduce this problem?

> 1. Set up the actuator using `http.Handle`
>
> ```go
> handler := actuator.GetHandler(&actuator.Config{})
> http.Handle("/actuator", handler)
> ```
>
> 2. Start the server
> 3. Send a request to the `/metrics` endpoint
>
> ```bash
> curl -X GET "http://localhost:3000/actuator/metrics"
> ```
>
> 4. The response contains the value 0x0583 for bysize

## Solution

Describe how you fixed the problem.
If you had assistance (i.e. Google search), provide links.

> A custom `Marshaler` function was introduced to serialize the
> `bySize` field. Marshaler functions are demonstrated in
> https://blergh.21er.org/26/Golang-Interfaces-and-Json-Marshaling

### Test Plan

Testing is an important part of demonstrating that your
bugfix works. Please provide a detailed walk-through about
how your test works, which files preform the test, and any
caveats and assumptions which were made during designing your test

### Implementation Risks

This section details any risks associated with merging in the
merge request.

## Approval Requirements

- [ ] Summary line provides ample details about what bug is being fixed
- [ ] Details section provides ample details about:
  - [ ] What is causing the bug
  - [ ] How the bug was identified
  - [ ] Exact steps to recreate the buggy behavior
- [ ] Solution section contains a detailed explanation on how the
      merge request will solve the behavior detailed in the
      _Details_ section
- [ ] Review Solution architecture if diagrams and drawings are
      provided
- [ ] Evidence (logs, screenshots, etc.) of the solution working
      is provided
- [ ] CI pipeline jobs pass
- [ ] Test coverage for the introduced code meets or exceeds 80%
