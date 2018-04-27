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
      {
         "Double Last Name",
         "%B0123456789012345^WAGNER JR/ELIZABETH A      ^" +
         "0123456U0000005591?;1234567890123456=01234567890123456?",
         "ELIZABETH",
         "WAGNER JR",
         "ELIZABETH WAGNER JR",
         "5591",
      },
      {
         "Hyphenated Last Name",
         "%B0123456789012345^TESTER-PERSON/JACOB        ^" +
         "0123456U0000002345?;1234567890123456=01234567890123456?",
         "JACOB",
         "TESTER-PERSON",
         "JACOB TESTER-PERSON",
         "2345",
      },
      {
         "Half-Swipe",
         "%B9876543212345678^WILLIAMS/GEORGE           ^" +
         "1234567U0000008888?",
         "GEORGE",
         "WILLIAMS",
         "GEORGE WILLIAMS",
         "8888",
      },
   }
   // Test Setup
   cardObj := New()

   for _, table := range tables {
      cardObj.Swipe(table.data)

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
            t.Errorf("ID generation was incorrect,\n\t " +
               "got: %q, want: %q.",
               cardObj.GetID(), table.id)
         }
      })
   }
}


