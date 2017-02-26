# ab2tar [![Build Status](https://travis-ci.org/bfritz/ab2tar.svg?branch=master)](https://travis-ci.org/bfritz/ab2tar)

Quick hack in Go to convert unencrypted Android backup files into
tar archives.  Strips off the magic bytes at the beginning and
inflates the rest of the zlib-compressed file.

Example usage:

    ab2tar -f com.android.providers.telephony.ab | tar -tvf -
