## PatientAnalyticsZKP

Our project aims to address the critical need for privacy in health care data without compromising on the ability to conduct vital research.
We are using Zero Knowledge Proofs and Pederson Commitment to encrypt patients information enabling users and researchers to execute analytical functions while preserving privacy.

## Summary of Key Components

# patient.go

Defines a structure for managing patient data while maintaining
privacy using zero-knowledge proofs (ZKPs) and Pedersen Commitments.

# ledger.go

Provides functionalities to initialize a ledger, add encrypted patient records to it, and filter patients based on their diseases while maintaining privacy through encryption and cryptographic operations.

# zkcircuit.go

Allows for the generation and verification of zero-knowledge proofs ensuring the equality of a patient's disease with a provided input disease, providing a privacy-preserving mechanism for verifying sensitive information.

# init.go

Sets up the cryptographic parameters required for the patient ledger system, enabling secure cryptographic operations such as point generation and manipulation on the curve.

```


# Run the following command for tests

```

go test

```
## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

Please make sure to update tests as appropriate.
```
