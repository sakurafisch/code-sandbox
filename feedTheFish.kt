fun swim(speed: String = "fast") {
    println("swimming $speed")
}

fun isTooHot(temperature: Int) : Boolean = temperature > 30
fun isTooDirty(dirty: Int) : Boolean = dirty > 30
fun isSunday(day: String) : Boolean = day == "Sunday"

fun shouldChangeWater (day: String, temperature: Int = 22, dirty: Int = 20): Boolean {
    return when {
        isTooHot(temperature) -> true
        isTooDirty(dirty) -> true
        isSunday(day) -> true
        else -> false
    }
}

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
    println("Change water: ${ shouldChangeWater(day) }")
}

fun main(args: Array<String>) {
    feedTheFish()
    swim()
    swim("slow")

}
