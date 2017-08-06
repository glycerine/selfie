selfie
========

A selfie is a useful indicator of identity.

A selfie is an SSH public-key that has been signed
by its own private key, together with a timestamp
generated at the point of signing.

If someone presents you with a selfie,
you can validate that the public key was signed
by the corresponding private key. This
signing happend at the time given in the selfie.

Selfies are much, much, easier to deal with
than the byzantine X.509 certificate stuff.
They don't expire, so your tests and servers won't
mysteriously stop working all of a sudden.
Signatures can be validated for recency.

Canonically, selfies are transferred using
`greenpack`[1]. For debugging (not transport),
they can be dumped to JSON. For example:

~~~
{"SignMe_zid00_rct":{"PubKeyBytes_zid00_bin":"AAAAB3NzaC1yc2EAAAADAQABAAAAgQDUXQteCgPhr52gpjqrRZ3n2sVyhxTL+2H3+vjnHn4F3e/8t6GQN7HSNT7ZUq/9DDCGnWOf7NIjAr2/wKRZ35n9K+9pQR5tclT1nBaFTXNf/eNwob9p29ZDkOoS0NQ9g3f4XNq/mdg/eCIRrUswAYo2FY88aTp9+sdzp2+syDP+1w==","SignedAtTimestamp_zid01_tim":"2017-08-06T18:33:49.016577586-05:00"},"SignatureFormat_zid01_str":"ssh-rsa","SignatureBlob_zid02_bin":"L2Di1NXBSlqTZJ+f+BAyGRX6ky1Pr0VLFQ2QR4OvVVfYcu5X8Fpx1GCUkTHFDiZbUN3BGpkAn44O07uTczkECyEcwtDzAJrTs4GXxxUS7ad4zVBs0BKMs/5KJ6eCFjvINgpH2kStivFdJv7YOYBvygjG5e9yrUGv2spKRE8gWL8="}
~~~

License: MIT

Author: Jason E. Aten, Ph.D.
