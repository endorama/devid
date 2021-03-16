# DevID

`devid` (pronounced `/ˈdeɪvɪd/`) is a Swiss Army Knife for your developer identities personas.

> A persona (plural personae or personas), depending to the context, can refer to either the public image of one's personality, or the social role that one adopts, or a fictional character.[1] The word derives from Latin, where it originally referred to a theatrical mask.[2] On the social web, users develop virtual personas as online identities. In fan fiction and in online stories, the personas may especially reflect the authors' self-insertion.

Each of us has multiple `personas` for different areas of their life. It may be work/personal, or for different open
source projects, or whatever reason you may think for presenting yourself differently in different context. This is
something we do in real life (think dressing differently for different social events) but doing so in digital
world as developers can be a pain: you have to manage identities (GPG or SSH keys), authentication tokens, specific
configurations.

Properly securing our developer identity and personas is hard. `devid` aims to help you with that.

> On the internet, nobody knows you're a dog

`devid` is not a tool for verifying identity (even if some day it may gain this feature). `devid` aims to easen the
managing or your personas. 1 or 100, it allows you to create and manage your identity consistently with safe and secure
defaults.

Security is hard and applying our collective intelligence can make it better.

# The history

I started working on `devenv` some years ago, as a way to manage different personas (at that time I called them profiles) while freelancing for different clients.

I maintained and used it over the years (starting around 2016 as a POC in BASH and evolving it in a Golang tool later on).

In 2020 I got in touch with [OpenSSF](https://openssf.org/) and joined the Digital Identity Attestation Working Group.  
During the meetings there have been a lot of discussion about "making security easy", and this is my take.

# The goal

This aims to be a comprehensive and extensible command line tool to manage your personas. The goal is not to reimplement
existing stuff but leverage awesome projects and **gluing them together applying secure defaults**.

Making security easy is **hard** but I think necessary to help adoption.

# The tool

## CLI

```
devid new <persona name> ✔
  Create a new persona configuration file, opens it within EDITOR
devid list ✔
  List all available personas
devid edit <persona name> ✔
  Open within EDITOR the specified persona configuration file
devid delete <persona name> ✔
  Delete a persona, securely removing all configurations
devid rehash <persona name>
  Recreate the configuration used by shell and run commands

devid shell --persona=<persona name>
  Spawn a shell with the specific persona environment loaded
devid run --persona=<persona name> <command>
  Run a one-off script in the specific persona environment

devid whoami ✔
  Print current persona name on stdout or exit with error if no persona is loaded

devid backup --persona=<persona name> ✔
  Backup the specified persona configurations & files
devid restore --persona=<persona name> --from-file=<backup file path>
  Restore the specified backup file in the specified persona

devid plugin list
  List available plugins
devid plugin active [--persona=<persona name>]
  List active plugins for current persona (or specified persona)

devid help ✔
  Print help text
devid version ✔
  Print cli version
```

# LICENSE

Apache License V2.0

# Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md)
