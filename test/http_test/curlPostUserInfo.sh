while true; do curl -XPOST  -d "firstnames=mamad" -d "lastnames=mammadi" -d "emails=mac@mac" -d "ticketnumbers=1" http://localhost/userinfos && sleep 3 && clear ; done

postUserInfo () { 
    while true do 
    curl -XPOST  -d "firstnames=mamad" -d "lastnames=mammadi" -d "emails=mac@mac" -d "ticketnumbers=1" http://localhost/userinfos && sleep 3 && clear 
    done
}

postUserInfo