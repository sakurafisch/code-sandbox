fun randomDay(): String {
    val week = arrayOf("Monday", "Tuesday", "Wednesday", "Thursday",
        "Friday", "Saturday", "Sunday")
    return week[java.util.Random().nextInt(week.size)]
}

fun fishFood(day: String) : String {
    return when (day) {
        "Monday" -> "flakes"
        "Tuesday" -> "redworms"
        "Wednesday" -> "granules"
        "Thursday" -> "granules"
        "Friday" -> "mosquitoes"
        "Saturday" -> "plankton"
        else -> "nothing"
    }
}

fun feedTheFish() {
    val day = randomDay()
    val food = fishFood(day)
    println("Today is $day and the fish eat $food")

}

fun main(args: Array<String>) {
    feedTheFish()
    feedTheFish()
    feedTheFish()
    feedTheFish()
    feedTheFish()
    feedTheFish()
}