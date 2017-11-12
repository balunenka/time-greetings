#!/bin/bash

    case $1 in 
        -d|--debug|-debug) 
                DEBUG_ENABLE="on"
                ;;
        -h|--help|-help)
                printf "TBD help text\n"
                ;;
        *) 
                DEBUG_ENABLE="off"
                ;;
esac




SCRIPT_NAME="$(basename "$(test -L "$0" && readlink "$0" || echo "$0")")"
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
REPO_DIR="$(git rev-parse --show-toplevel)"
BIN_DIR="${REPO_DIR}/src/bin"



source ${BIN_DIR}/colors.sh
source ${BIN_DIR}/DEBUG-enable.sh ${DEBUG_ENABLE}
source ${BIN_DIR}/goos_declare.sh




DEBUG printf "${RED}DEBUG: ${NORMAL}Path of script ${SCRIPT_NAME} location = ${SCRIPT_DIR}\n"
DEBUG printf "${RED}DEBUG: ${NORMAL}Path of bin location = ${BIN_DIR}\n"



##########################################################################################################
###   Set PACKAGE_NAME as directory of package (script should be run from this directpry)              ###
##########################################################################################################
function get-package-name(){

    PACKAGE_NAME=$(basename "$SCRIPT_DIR")
    DEBUG printf "${RED}DEBUG: ${NORMAL}Package name is ${YELLOW}${PACKAGE_NAME}${NORMAL}\n"    

}


function zip-executable(){
    
    DEBUG printf "${RED}DEBUG: ${NORMAL}Put all executable of${YELLOW}${PACKAGE_NAME}${NORMAL} to archive ${PACKAGE_NAME}.zip ....\n"

    zip ${PACKAGE_NAME}.zip ${PACKAGE_NAME}*

    DEBUG printf "${RED}DEBUG: ${NORMAL}Archiving is done\n"

    DEBUG print-separator-line

}


function remove-executable(){

    DEBUG printf "${RED}DEBUG: ${NORMAL}Delete all previous builds....\n"
    DEBUG ls -l |grep "${PACKAGE_NAME}-"
    rm ${PACKAGE_NAME}-*
    DEBUG printf "${RED}DEBUG: ${NORMAL}All previous builds deleted\n"
    DEBUG ls -l |grep "${PACKAGE_NAME}-"
    DEBUG print-separator-line


}




##########################################################################################################
###   Builds executables for all platforms that stores in GOOS_FILE (value get frogoos_declare.sh)     ###
###   function check-goos-file-exist() get from src/bin/goos_declare.sh                                ###
###   function print-separator-line() get from src/bin/DEBUG-enable.sh                                 ###
##########################################################################################################
function go-build(){
   
    get-package-name
    remove-executable
   

    
    check-goos-file-exist
    DEBUG printf "GOOS file in ${SCRIPT_NAME} is ${GOOS_FILE} \n\n"

    
        while read -r line
        do
            GOOS="$line"
            DEBUG printf "${RED}DEBUG: ${NORMAL}Name read from file - $GOOS"
            DEBUG printf "${RED}DEBUG: ${NORMAL}Build ${YELLOW} ${PACKAGE_NAME} ${NORMAL} for ${GOOS} platform\n"
            go build
            DEBUG printf "${RED}DEBUG: ${NORMAL}Build done\n"

            GO_EXECUTABLE="${PACKAGE_NAME}-${GOOS}"
            case $GOOS in
                windows) 
                    DEBUG printf "${RED}DEBUG: ${NORMAL}Renaming executable to ${GO_EXECUTABLE}......\n"
                    mv ${PACKAGE_NAME}.exe ${GO_EXECUTABLE}.exe
                    DEBUG printf "${RED}DEBUG: ${NORMAL}Rename done\n"
                    DEBUG print-separator-line
                    ;;
                *)
                    DEBUG printf "${RED}DEBUG: ${NORMAL}Renaming executable to ${GO_EXECUTABLE}......\n"
                    mv ${PACKAGE_NAME} ${GO_EXECUTABLE}
                    DEBUG printf "${RED}DEBUG: ${NORMAL}Rename done\n"

                    DEBUG print-separator-line
                    ;;
            esac

        done < "$GOOS_FILE"

        
   
}




# Run build
go-build

#zip-executable






