#!/bin/bash

if [[ $# -ne 2 ]]; then
    echo "expected 2 arguments got $#"
    echo "please provide {path_to_version.txt} (major|minor|patch)"
    exit 1
fi

version_file_path=$1
type=$2

if [[ $type != "major" && $type != "minor" && $type != "patch" ]]; then
    echo "please provide (major|minor|patch)"
    exit 1
fi

version=$(cat $version_file_path)
if [[ $? -ne 0 ]]; then
    echo "failed to read version file"
    exit 1
fi


# removes v to make it easier to parse
v_removed=$(echo $version | sed -E 's/v//')

# accepts (major|minor|patch)
major_version=$(echo $v_removed | cut -d'.' -f1)
minor_version=$(echo $v_removed | cut -d'.' -f2)
patch_version=$(echo $v_removed | cut -d'.' -f3)

if [[ $type == "major" ]]; then
    major_version=$(($major_version + 1))
elif [[ $type == "minor" ]];then
    minor_version=$(($minor_version + 1))
elif [[ $type == "patch" ]];then
    patch_version=$(($patch_version + 1))
fi

final_version="v$major_version.$minor_version.$patch_version"
echo "version bump: $version -> $final_version"
echo $final_version > $version_file_path