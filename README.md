# secret-santa-generator
Generate secret santa assignments for your family and friends


# Usage

1. Add all secret santa participants to `config.json` with a value for each field

    See `example.config.json` for format, convert with:
    ```
    mv ./example.config.json ./config.json
    ```

2. Set auth environment variables
    - `HERMES_SENDER_EMAIL`
    - `HERMES_SENDER_PASSWORD`

3. [Allow user/password authentication](https://myaccount.google.com/lesssecureapps) on your Gmail account 
   
    **Note-** This makes your gmail account less secure and should be re-enabled after you've run the tool.

4. Run the generator

    ```
    go run gen.go
    ```

