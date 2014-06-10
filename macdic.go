package main

import (
  "os"
  "unsafe"
)

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa,CoreFoundation

#import <Foundation/Foundation.h>
#import <Cocoa/Cocoa.h>
#import <COreServices/CoreServices.h>
#include <stdio.h> // FIXME: delete

NSArray *DCSGetActiveDictionaries();
NSArray *DCSCopyAvailableDictionaries();
NSString *DCSDictionaryGetName(DCSDictionaryRef dictID);
NSString *DCSDictionaryGetShortName(DCSDictionaryRef dictID);

void lookup(const char *dicShortName, const char *word) {
  id pool = [NSAutoreleasePool new];

  NSString *nsDicShortName = [NSString stringWithUTF8String:dicShortName];
  NSString *nsWord = [NSString stringWithUTF8String:word];

  CFRange range;
  range.location = 0;
  range.length = [nsWord length];

  DCSDictionaryRef dic = NULL;
  NSArray *dicts = DCSCopyAvailableDictionaries();
  for (NSObject *aDict in dicts) {
    NSString *aShortName = DCSDictionaryGetShortName((DCSDictionaryRef)aDict);
    if ([aShortName isEqualToString:nsDicShortName]) {
      dic = (DCSDictionaryRef)aDict;
    }
  }

  CFStringRef ref = NULL;
  ref = DCSCopyTextDefinition(dic, (CFStringRef)nsWord, range);

  printf("%s\n", [(NSString*)ref UTF8String]);

  [pool drain];
}
*/
import "C"

func main() {
  if len(os.Args) <= 1 {
    os.Exit(1)
  }
  dicShortName := C.CString("Japanese-English")
  defer C.free(unsafe.Pointer(dicShortName))
  word := C.CString(os.Args[1])
  defer C.free(unsafe.Pointer(word))
  C.lookup(dicShortName, word)
}
