"""
Advent of Code 2022 Day 6

Date: 12/6/2022
Author: phortheman

"""
from pathlib import Path

PACKET_SIZE = 4
MESSAGE_SIZE = 14

with open( Path(__file__).with_name( "input.txt" ), 'r' ) as file:
    datastream = file.readline()
    startOfPacketMarker = 0
    startOfMessageMarker = 0
    for i in range( len(datastream ) ):
        if( startOfPacketMarker == 0 and len(set(datastream[i:i+PACKET_SIZE] ) ) == PACKET_SIZE ):
            startOfPacketMarker = i + PACKET_SIZE
        
        if( startOfMessageMarker == 0 and len(set(datastream[i:i+MESSAGE_SIZE] ) ) == MESSAGE_SIZE ):
            startOfMessageMarker = i + MESSAGE_SIZE
        
        if( startOfPacketMarker != 0 and startOfMessageMarker != 0 ):
            break
    
    print( f"Start-of-packet marker: {startOfPacketMarker}" )
    print( f"Start-of-message marker: {startOfMessageMarker}" )
