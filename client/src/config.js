export const BASE_API_URL = process.env.BASE_API_URL || "http://localhost:4000/api";
export const ENUM_MAPPINGS = {
    "virtualMachineState": {
        "0": "Stopped",
        "1": "Running",
        "2": "Error",
    },
    "months": {
        1: "January",
        2: "February",
        3: "March",
        4: "April",
        5: "May",
        6: "June",
        7: "July",
        8: "August",
        9: "September",
        10: "October",
        11: "November",
        12: "December",
    }
};
export const DATA_DEFAULTS = {
    "virtualMachine": {
        "cores": [1, 2, 4],
        "diskSpace":  [128, 256, 512, 1024],
        "ram": [1, 4, 8, 16],
    }
}