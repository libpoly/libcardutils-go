package cardutils

import (
   "testing"
)

func TestCardDataGeneration(t *testing.T) {
   tables := []struct{
      testCase string
      data string
      fName string
      lName string
      fullName string
      id string
   }{
      {
         "Short Name",
         "%B0123456789012345^SMITH/JOHN Z              ^" +
         "0123456U0000008263?;1234567890123456=01234567890123456?",
         "JOHN",
         "SMITH",
         "JOHN SMITH",
         "8263",
      },
      {
         "Long Name",
         "%B0123456789012345^ANDERSON/ALEXANDER A      ^" +
         "0123456U0000004290?;1234567890123456=01234567890123456?",
         "ALEXANDER",
         "ANDERSON",
         "ALEXANDER ANDERSON",
         "4290",
      },
   }

   for _, table := range tables {
      // Test Setup
      cardObj := CreateCard(table.data)

      // Test Execution
      t.Run(table.testCase, func(t *testing.T) {
         if cardObj.GetFirstName() != table.fName {
            t.Errorf("First name generation " +
               "was incorrect,\n\tgot: %q, want: %q.",
               cardObj.GetFirstName(), table.fName)
         }
         if cardObj.GetLastName() != table.lName {
            t.Errorf("Last name generation " +
               "was incorrect,\n\tgot: %q, want: %q.",
               cardObj.GetLastName(), table.lName)
         }
         if cardObj.GetFullName() != table.fullName {
            t.Errorf("Full name generation " +
               "was incorrect,\n\tgot: %q, want: %q.",
               cardObj.GetFullName(), table.fullName)
         }
         if cardObj.GetID() != table.id {
            t.Errorf("ID generation was incorrect,\n\t got: %q, want: %q.",
               cardObj.GetID(), table.id)
         }
      })
   }
}


