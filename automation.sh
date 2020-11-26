#!/bin/bash

action=$1
chapter=$2
workdir=$(echo $MLGO)
target=()
cnt=0

check_files() {
    cd ${workdir}
    dirs=$(ls -d */ | grep 'ch')

    # chapters
    for dir in ${dirs}; do
        ch="${workdir}/${dir}"
        cd ${ch}
        subs=$(ls -d */)

        # subjects
        for sub in ${subs}; do
            s="${ch}${sub}"
            cd ${s}

            if [[ "$action" = "build" ]]
            then
                files=$(ls -d -- [0-9][0-9]*.go)

            elif [[ "$action" = "clean" || "$action" = "run" ]]
            then
                files=$(ls -d -- [0-9][0-9]* | grep -v '.\.go')
            fi
            
            # files
            for f in ${files}; do
                target+=(${s}${f})
            done

            cd ${ch}
        done

        cd ${workdir}
    done
}

build() {
    for file in ${target[@]}; do
        if [ -f ${file} ]
        then
            # https://tldp.org/LDP/LG/issue18/bash.html
            # directory
            cd ${file%/*}

            if [[ -z "$chapter" ]]
            then 
                echo "Build ${file}"
                # file
                go build ${file##*/}
                ((cnt++))
            elif [[ $file == *"$chapter"* ]]
            then
                echo "Build ${file}"
                go build ${file##*/}
                ((cnt++))
            fi
        fi
    done
}

clean() {
    for file in ${target[@]}; do
        if [ -f ${file} ]
        then
            if [[ -z "$chapter" ]]
            then 
                echo "Remove ${file}"
                rm ${file}
                ((cnt++))
            elif [[ $file == *"$chapter"* ]]
            then
                echo "Remove ${file}"
                rm ${file}
                ((cnt++))
            fi
        fi
    done
}

run() {
    for file in ${target[@]}; do
        if [ -f ${file} ]
        then
            # directory
            cd ${file%/*}

            if [[ -z "$chapter" ]]
            then 
                echo "Runned ${file}"
                # file
                ./${file##*/}
                ((cnt++))
            elif [[ $file == *"$chapter"* ]]
            then
                echo "Runned ${file}"
                ./${file##*/}
                ((cnt++))
            fi
        fi
    done
}

check_files

if [[ "$action" = "build" ]]
then
    build
    echo; echo "Total built: $cnt"
elif [[ "$action" = "clean" ]]
then
    clean
    echo; echo "Total cleaned: $cnt"
elif [[ "$action" = "run" ]]
then
    run
    echo; echo "Total runuted: $cnt"
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
