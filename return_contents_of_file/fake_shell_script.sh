#!/bin/bash


################################################################################
#~~~~~~~~~~~~~~~~~~~~~~~ DETAILS ABOUT THE SHELL SCRIPT ~~~~~~~~~~~~~~~~~~~~~~~#
################################################################################
#                                                                              #
#                                                                              #
################################################################################


################################################################################
#~~~~~~~~~~~~~~~~~~~~~~~~~~~~ SET SCRIPT VARIABLES ~~~~~~~~~~~~~~~~~~~~~~~~~~~~#
################################################################################
#******************************************************************#
# The job name to execute. We will assume a few things with this:  #
#                                                                  #
#   1) Unless overridden, the full job name will be derived by     #
#      taking the ${JOB_NAME} and replacing the dashes with spaces #
#      and making the first letter of each word capitalized.       #
#                                                                  #
#   2) The job specific config and OCP config will reside in the   #
#      same directory as this script file and will be named as     #
#      "${JOB_NAME}-config.xml".                                   #
#                                                                  #
#   3) The job specific log file will reside in "../logs/jobs/"    #
#      and will be named as "${JOB_NAME}.log".                     #
#                                                                  #
#   If spring batch job...                                         #
#   4) The spring batch bean configuration in the classpath file   #
#      will reside in "META-INF/spring/batch/" and will be named   #
#      as "${JOB_NAME}-batch-job.xml". This file must also contain #
#      a batch:job node with it's id attribute equal to the job    #
#      name specified below.                                       #
#******************************************************************#
export JOB_NAME="shell_script_fake"


################################################################################
#~~~~~~~~~~~~~~~~~~~~~~ SET PRE AND POST COMMAND SCRIPTS ~~~~~~~~~~~~~~~~~~~~~~#
################################################################################
#******************************************************************#
# Specify the script that you would like to run prior to the       #
# command execution. Specifying it here will run the script within #
# the lock file creation and deletion, preventing multiple         #
# executions of this script running at the same time. This value   #
# can be left blank or section can be removed, if not needed.      #
#******************************************************************#
export FAKE_COMMAND_SCRIPT=

#******************************************************************#
# Specify the script that you would like to run after the          #
# command execution. Specifying it here will run the script within #
# the lock file creation and deletion, preventing multiple         #
# executions of this script running at the same time. This value   #
# can be left blank or section can be removed, if not needed.      #
#******************************************************************#
export FAKE_COMMAND_SCRIPT=

#******************************************************************#
# Setting this flag to Y will prevent subsequent scripts from      #
# running if the previous script returns a value other than zero.  #
# If the pre-command script fails then the main command and the    #
# post-command will not execute.                                   #
#******************************************************************#
export STOP_ON_FAILURE=N


################################################################################
#~~~~~~~~~~~~ SET ENVIRONMENT, CLASSPATH, ETC. AND EXECUTE COMMAND ~~~~~~~~~~~~#
################################################################################
