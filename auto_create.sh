#!/usr/bin/env bash

function _create_folder() {
	title=$1
	fileSuffix=$2
	title_path=$3

	FileArray=("${title}"."${fileSuffix}" README.md)
    echo "Create ${title} folder: ${title_path}"
    for i in "${FileArray[@]}"; do
        i_path=${title_path}/${i}
        echo "filepath: ${i_path}"
        mkdir -p "${title_path}"
        if [[ ! -e ${i_path} ]]; then
            touch "${i_path}"
        fi
    done
}


function CreateFolder() {
    echo "Input Python or Go(Case-Insensitive) or All(Include both of go and python):"
    read -r program_name
    echo "Input the title of leetcode, which will be the sub-folder name:"
    read -r title

    case ${program_name} in
    [Pp][Yy][Tt][Hh][Oo][Nn] | [Pp][Yy])
        fileSuffix="py"
        titlePath=./leetcode/Python/code/${title}
        ;;

    [Gg][Oo])
        fileSuffix="go"
        titlePath=./leetcode/Go/code/${title}
        ;;
	[Aa][Ll][Ll])
		echo "************* Py *************"
		_create_folder "$title" "py" "./leetcode/Python/code/${title}"
		echo "************* Go *************"
		_create_folder "$title" "go" "./leetcode/Go/code/${title}"
		echo "Task done!!!"
		exit 1
		;;
    *)
        echo "Input error, please try again."
        exit 1
        ;;
    esac

	_create_folder "$title" "$fileSuffix" "$titlePath"
    
}

CreateFolder
echo "Task done!!!"
