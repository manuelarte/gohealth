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

Complete your bug request by making sure the following is provided:

- [ ] Attach screenshots which demonstrate the buggy behavior
- [ ] Provide the version number for which you observed the buggy behavior
- [ ] Provide OS and Go Runtime information that you used with
      the library
- [ ] If possible, attach relevant logs related to the buggy behavior
