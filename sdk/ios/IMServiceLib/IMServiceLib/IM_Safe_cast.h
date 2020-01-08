//
//  Safe_cast.h
//  testImage
//
//  Created by li on 10/8/15.
//  Copyright © 2015 hh.changhong.com. All rights reserved.
//

#import <Foundation/Foundation.h>

@interface IM_Safe_cast : NSObject


+ (NSString *)parseToString:(id)parse;

+ (float)parseToFloatValue:(id)parse;

+ (int)parseToIntValue:(id)parse;
+ (long)parseTolongValue:(id)parse;
+ (long long)parseTolongLongValue:(id)parse;
+ (unsigned long long)parseToUnsignedLongLongValue:(id)parse;

+ (NSInteger)parseToIntegerValue:(id)parse;

+ (BOOL)parseToBOOLValue:(id)parse;
+ (double)parseToDoubleValue:(id)parse;
+ (NSNumber *)parseToNumValue:(id)parse;
+ (NSArray *)parseToArray:(id)parse;
+ (NSDictionary *)parseToDictionary:(id)parse;


@end
