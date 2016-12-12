class Colors(object):
    BLUE = "\033[94m"
    GREEN = "\033[92m"
    BOLD = "\033[1m"
    COLOREND = "\033[0m"
    FAIL = "\033[91m"
    PURPLE = "\033[95m"


building_msg = Colors.BOLD + Colors.GREEN +  "[BUILDING]:   " + Colors.COLOREND
cleaning_msg = Colors.BOLD + Colors.PURPLE + "[CLEANING]:   " + Colors.COLOREND
install_msg  = Colors.BOLD + Colors.BLUE +   "[INSTALLING]: " + Colors.COLOREND
running_msg  = Colors.BOLD + Colors.BLUE +   "[RUNNING]:    " + Colors.COLOREND
