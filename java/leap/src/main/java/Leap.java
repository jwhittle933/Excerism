class Leap {

    boolean isLeapYear(int year) {
        boolean divByFour = year%4 == 0;
        boolean divByHundred = year%100 == 0;
        boolean divByFourHundred = year%400 == 0;

        if (divByFour) {
            if (divByHundred) {
                if  (divByFourHundred) {
                    return true;
                }
                return false;
            }
            return true;
        }
        return false;
    }

}
