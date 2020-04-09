commp2cid
=========

Conversion tool from Piece commitment hashes used by [Filecoin] to Content IDs
that meet [the CIDv1 standard].

## Install

There are no pre-built binaries, it has to be built from source.

The Go toolchain available must be available in the system. For development
convenience this project uses [asdf-vm] with the respective [asdf-golang]
plugin.

```shell
git clone https://github.com/subvisual/commp2cid.git --branch v1.0.1
cd commp2cid
make
```

For system-wide installation, move or link the `commp2cid` executable into a
location in the system's `PATH`.

## Usage

```shell
commp2cid < commps.txt
```

This tool maps each line of the standard input as a commitment hash into a line
with the respective CID in the standard output.

### Caveats

- It **does not** read from a file;
- It **does not** process any positional arguments;
- It **does not** validate the provided commitment hashes.

## Background

See the [Filecoin PoRep Spec] and the [Filecoin Paper] for how these piece
commitment hashes are generated.

The conversion is handled by [filecoin-project/go-fil-commcid].

## License

This repository is under the [MIT license].

Copyright 2020. [Subvisual, Lda].

[Filecoin]: https://filecoin.io
[the CIDv1 standard]: https://github.com/multiformats/cid
[asdf-vm]: https://asdf-vm.com
[asdf-golang]: https://github.com/kennyp/asdf-golang
[Filecoin Paper]: https://filecoin.io/filecoin.pdf
[Filecoin PoRep Spec]: https://filecoin-project.github.io/specs/#algorithms__porep
[filecoin-project/go-fil-commcid]: https://github.com/filecoin-project/go-fil-commcid
[MIT license]: ./LICENSE.txt
[Subvisual, Lda]: https://subvisual.com
