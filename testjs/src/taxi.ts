export const Taxi = (kilometer: number, minute: number) => {
    const meter = kilometer * 100
    const roundedKilometer = meter % 50 == 0 ? kilometer : (meter + (50 - (meter % 50))) / 100
    const second = minute * 60
    const roundedMinute = second % 60 == 0 ? minute : (second + (60 - (second % 60))) / 60
    //const result = (4 * roundedKilometer) + roundedMinute
    return  (4 * roundedKilometer) + roundedMinute
}
