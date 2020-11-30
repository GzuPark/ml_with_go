#!/bin/bash

ACTION=$1
CHAPTER=$2
WORKDIR=$(echo $MLGO)
TARGET=()
CNT=0

check_files() {
    cd ${WORKDIR}
    dirs=$(ls -d */ | grep 'ch')

    # chapters
    for dir in ${dirs}; do
        ch="${WORKDIR}/${dir}"
        cd ${ch}
        subs=$(ls -d */)

        # subjects
        for sub in ${subs}; do
            s="${ch}${sub}"
            cd ${s}

            if [[ "$ACTION" = "build" ]]
            then
                files=$(ls -d -- [0-9][0-9]*.go)

            elif [[ "$ACTION" = "clean" || "$ACTION" = "run" ]]
            then
                files=$(ls -d -- [0-9][0-9]* | grep -v '.\.go' | grep -v '.\.md')
            fi
            
            # files
            for f in ${files}; do
                if [[ ${s} != *"ch09"* ]]
                then
                    TARGET+=(${s}${f})
                fi
            done

            cd ${ch}
        done

        cd ${WORKDIR}
    done
}

build() {
    for file in ${TARGET[@]}; do
        if [ -f ${file} ]
        then
            # https://tldp.org/LDP/LG/issue18/bash.html
            # directory
            cd ${file%/*}

            if [[ -z "$CHAPTER" ]]
            then 
                echo "Build ${file}"
                # file
                go build ${file##*/}
                ((CNT++))
            elif [[ $file == *"$CHAPTER"* ]]
            then
                echo "Build ${file}"
                go build ${file##*/}
                ((CNT++))
            fi
        fi
    done
}

clean() {
    for file in ${TARGET[@]}; do
        if [ -f ${file} ]
        then
            if [[ -z "$CHAPTER" ]]
            then 
                echo "Remove ${file}"
                rm ${file}
                ((CNT++))
            elif [[ $file == *"$CHAPTER"* ]]
            then
                echo "Remove ${file}"
                rm ${file}
                ((CNT++))
            fi
        fi
    done
}

run() {
    for file in ${TARGET[@]}; do
        if [ -f ${file} ]
        then
            # directory
            cd ${file%/*}

            if [[ -z "$CHAPTER" ]]
            then 
                echo "Runned ${file}"
                # file
                ./${file##*/}
                ((CNT++))
            elif [[ $file == *"$CHAPTER"* ]]
            then
                echo "Runned ${file}"
                ./${file##*/}
                ((CNT++))
            fi
        fi
    done
}

check_files

if [[ "$ACTION" = "build" ]]
then
    build
    echo; echo "Total built: $CNT"
elif [[ "$ACTION" = "clean" ]]
then
    clean
    echo; echo "Total cleaned: $CNT"
elif [[ "$ACTION" = "run" ]]
then
    run
    echo; echo "Total runuted: $CNT"
else
    echo; echo "Assign 1st argument: ( build || clean || run )"
    echo; echo "Example:"
    echo "    \$MLGO/run.sh build"
    echo "    \$MLGO/run.sh clean"
    echo "    \$MLGO/run.sh run"
    echo "    \$MLGO/run.sh build ch01"
    echo "    \$MLGO/run.sh clean ch01"
    echo "    \$MLGO/run.sh run ch01"; echo
fi
