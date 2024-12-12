# Class for a Bingo Card Field

# A BingoField is a data struct that have the value of the field and if it is was pulled
class BingoField:
    def __init__(self,value) -> None:
        self.value = value
        self.match = False

    def checkIfMatch( self, checkValue ):
        if( self.value == checkValue ):
            self.match = True

    def __repr__(self) -> str:
        return str(self.value) + ":" + str(self.match)

    def __str__(self) -> str:
        return str(self.value) + ":" + str(self.match)

# A BingoBoard is a data struct that contains a 2D list of BingoFields 
# And contains a dictionary of all the values and their coordinates on the 2D list
class BingoBoard:
    def __init__( self, twoDimListOfFields ):
        self.valueIndexes = { } # Key: value of the field, Value: "x,y" coord on the board
        self.board = twoDimListOfFields.copy()

        for xIndex, xItem in enumerate( self.board ):
            for yIndex, yItem in  enumerate( xItem ):
                self.valueIndexes[yItem.value] = f"{xIndex},{yIndex}"

    def __repr__(self):
        return self.board

    def checkBoardForValue( self, drawnValue ):
        
        onBoard = self.valueIndexes.get(drawnValue)
        if ( onBoard != None ):
            [row, col] = [ int(i) for i in onBoard.split(',') ]
            field = self.board[row][col]
            field.match = True
            return True
        else: 
            return False
    
    def checkIfBoardWon(self):
        victoryCheck = 0
        for row in self.board:
            for field in row:
                if( field.match ):
                    victoryCheck += 1
                    if( victoryCheck == 5 ):
                        return True
                else: 
                    victoryCheck = 0
                    break
        
        transposedBoard = [list(i) for i in zip(*self.board)]
        for col in transposedBoard:
            for field in col:
                if( field.match ):
                    victoryCheck += 1
                    if( victoryCheck == 5 ):
                        return True
                else:
                    victoryCheck = 0
                    break

        return False

    def sumOfUnmatchedFields( self ):
        sum = 0
        for row in self.board:
            for field in row:
                if( not field.match ):
                    sum += int(field.value)
        return sum
